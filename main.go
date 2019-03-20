package main

import (
	"context"
	"crypto/sha256"
	"net"
	"os"

	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
	"github.com/namsral/flag"
	"github.com/pkg/errors"
)

type imageMap struct {
	sync.RWMutex
	m map[string][]byte
}

func NewImageMap() *imageMap {
	return &imageMap{
		m: map[string][]byte{},
	}
}

func (im *imageMap) Add(h string, d []byte) {
	im.Lock()
	im.m[h] = d
	im.Unlock()
	logrus.Debugf("added %s", h)
}

func (im *imageMap) Remove(h string) {
	im.Lock()
	delete(im.m, h)
	im.Unlock()
	logrus.Debugf("removed %s", h)
}

func (im *imageMap) Get(h string) ([]byte, bool) {
	im.RLock()
	bytes, ok := im.m[h]
	im.RUnlock()
	return bytes, ok
}

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "SVG2PNG", 0)
	flagPort := fs.Int("port", 8544, "port to listen to")
	flagTimeout := fs.Int("timeout", 30, "initial timeout")
	flagURLs := fs.String("urls", "", "urls to chrome rdp (csv)")
	flagHosts := fs.String("hosts", "", "hosts with running chrome rdp (csv)")
	flagSelf := fs.String("self", "svg2png", "url under which chrome can reach this service (port is added automatically)")
	fs.Parse(os.Args[1:])

	if *flagHosts == "" && *flagURLs == "" {
		*flagURLs = "http://localhost:9222/json"
	}

	selfURL := fmt.Sprintf("http://%s:%d/v1/svg-html/", *flagSelf, *flagPort)
	logrus.SetLevel(logrus.DebugLevel)
	chromes, err := createCDPClients(*flagURLs, *flagHosts, *flagTimeout)
	if err != nil {
		logrus.Fatal(err)
	}
	images := NewImageMap()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/svg-html/", htmlHandler)
	mux.HandleFunc("/v1/svg-data/", dataHandler(images))
	mux.HandleFunc("/v1/png", mainHandler(images, chromes, selfURL))
	mux.HandleFunc("/healthz", healthzHandler)

	logrus.Debugf("listening on :%d", *flagPort)
	http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), mux)
}

func fetchImages(url *url.URL, res *[]byte) chromedp.Tasks {
	sel := `#svg`
	return chromedp.Tasks{
		chromedp.Navigate(url.String()),
		//chromedp.Sleep(2000 * time.Millisecond),
		chromedp.WaitVisible(sel, chromedp.ByID),
		//chromedp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, chromedp.ByQuery),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
		//chromedp.CaptureScreenshot(res),
	}
}

func createCDPClients(url, host string, timeout int) (chan *chromedp.CDP, error) {
	if url != "" && host != "" {
		return nil, fmt.Errorf("url and host parameters are mutually exclusive(u:'%s', h:'%s'", url, host)
	}
	var urls []string
	switch {
	case host != "":
		urls = make([]string, 0, 5)
		for _, h := range strings.Split(host, ",") {
			addrs, err := net.LookupHost(h)
			if err != nil {
				return nil, errors.Wrapf(err, "could not resolve '%s'", h)
			}
			for _, a := range addrs {
				urls = append(urls, fmt.Sprintf("http://%s:9222/json", a))
			}
		}
	default: // url parameter has a default so will never be empty
		urls = strings.Split(url, ",")
	}
	if len(urls) == 0 {
		return nil, errors.New("at least one chrome instance must be reachable")
	}

	s := "s"
	if len(urls) == 1 {
		s = ""
	}
	logrus.Infof("using %d chrome instance%s as target%s", len(urls), s, s)

	ctx := context.Background()

	chromes := make(chan *chromedp.CDP, len(urls))
	for _, u := range urls {
		u = strings.TrimSpace(u)
		start := time.Now()
		var c *chromedp.CDP
		var err error
		for {
			c, err = chromedp.New(ctx,
				chromedp.WithTargets(client.New(client.URL(u)).WatchPageTargets(ctx)),
				chromedp.WithErrorf(logrus.Errorf))
			if err == nil ||
				int(time.Now().Sub(start).Seconds()) > timeout {
				break
			}
			logrus.Debugf("trying '%s' again after %s", u, err)
		}
		if err != nil {
			return nil, err
		}
		chromes <- c
	}

	return chromes, nil
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	ch := r.URL.Path[len("/v1/svg-html/"):]
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(`<html><body><img id="svg" src="/v1/svg-data/` + ch + `" /></body></html>`))
}

func dataHandler(images *imageMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch := r.URL.Path[len("/v1/svg-data/"):]

		bytes, ok := images.Get(ch)
		if !ok {
			logrus.Warn(ch)
			http.Error(w, "404 Image not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		bw, err := w.Write(bytes)
		if err != nil {
			logrus.Warn(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if bw != len(bytes) {
			logrus.Warnf("incomplete write %s: %d/%d", ch, bw, len(bytes))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func mainHandler(images *imageMap, chromes chan *chromedp.CDP, selfURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := sha256.New()
		h.Write([]byte(time.Now().UTC().String()))
		body, err := ioutil.ReadAll(io.TeeReader(r.Body, h))
		if err != nil {
			logrus.Warn(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ch := fmt.Sprintf("%x.svg", h.Sum([]byte{}))
		images.Add(ch, body)
		defer images.Remove(ch)
		imageURL, err := url.Parse(fmt.Sprintf("%s%s", selfURL, ch))
		if err != nil {
			logrus.Warn(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var res []byte
		c := <-chromes
		err = c.Run(r.Context(), fetchImages(imageURL, &res))
		chromes <- c
		if err != nil {
			logrus.Warn(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(res)
	}
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK\n"))
}
