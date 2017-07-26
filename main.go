package main

import (
	"context"
	"crypto/sha256"
	"log"
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
	cdp "github.com/knq/chromedp"
	"github.com/knq/chromedp/client"
	"github.com/namsral/flag"
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
	flagURLs := fs.String("url", "http://localhost:9222/json", "urls to chrome (csv)")
	fs.Parse(os.Args[1:])

	logrus.SetLevel(logrus.DebugLevel)

	chromes := createCDPClients(*flagURLs, *flagTimeout)
	images := NewImageMap()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/svg-html/", htmlHandler)
	mux.HandleFunc("/v1/svg-data/", dataHandler(images))
	mux.HandleFunc("/v1/png", mainHandler(images, chromes))

	logrus.Debugf("listening on :%d", *flagPort)
	http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), mux)
}

func fetchImages(url *url.URL, res *[]byte) cdp.Tasks {
	sel := `#svg`
	return cdp.Tasks{
		cdp.Navigate(url.String()),
		//cdp.Sleep(2000 * time.Millisecond),
		cdp.WaitVisible(sel, cdp.ByID),
		//cdp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, cdp.ByQuery),
		cdp.Screenshot(sel, res, cdp.NodeVisible, cdp.ByID),
		//cdp.CaptureScreenshot(res),
	}
}

func createCDPClients(url string, timeout int) chan *cdp.CDP {
	urls := strings.Split(url, ",")
	s := "s"
	if len(urls) == 1 {
		s = ""
	}
	logrus.Infof("using %d chrome instance%s as target%s", len(urls), s, s)

	chromes := make(chan *cdp.CDP, len(urls))
	for _, u := range urls {
		u = strings.TrimSpace(u)
		start := time.Now()
		var c *cdp.CDP
		var err error
		for {
			c, err = cdp.New(context.Background(),
				cdp.WithTargets(client.New(client.URL(u)).WatchPageTargets(nil)),
				cdp.WithErrorf(logrus.Errorf))
			if err == nil ||
				int(time.Now().Sub(start).Seconds()) > timeout {
				break
			}
			logrus.Debugf("trying '%s' again after %s", u, err)
		}
		if err != nil {
			log.Fatalf("%s", err)
		}
		chromes <- c
	}

	return chromes
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

func mainHandler(images *imageMap, chromes chan *cdp.CDP) http.HandlerFunc {
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
		imageURL, err := url.Parse(fmt.Sprintf("http://svg2png:8544/v1/svg-html/%s", ch))
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
