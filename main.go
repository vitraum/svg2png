package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/Sirupsen/logrus"
	cdp "github.com/knq/chromedp"
)

func main() {
	flagPort := flag.Int("port", 8544, "port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/png", func(w http.ResponseWriter, r *http.Request) {
		tmpDir, err := ioutil.TempDir("", "svg2png")
		if err != nil {
			logrus.Warn(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() {
			rErr := os.RemoveAll(tmpDir)
			if rErr != nil {
				logrus.Warningf("Could not remove calculation tempdir: %s: %s", tmpDir, rErr)
			}
		}()

		reader, err := r.MultipartReader()
		if err != nil {
			logrus.Warn(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		i := 0
		images := []string{}
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			fname := path.Join(tmpDir, fmt.Sprintf("%d.svg", i))
			dst, err := os.Create(fname)
			defer dst.Close()
			if err != nil {
				logrus.Warn(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if _, err := io.Copy(dst, part); err != nil {
				logrus.Warn(err)

				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			i++
			images = append(images, fname)
		}

		c, err := cdp.New(r.Context(), cdp.WithLog(log.Printf))
		if err != nil {
			log.Fatal(err)
		}

		// run task list
		var res = [][]byte{}
		err = c.Run(r.Context(), fetchImages(images, &res))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("saved screenshot of #testimonials from search result listing `%s` (%s)", res, site)

		w.Write([]byte(fmt.Sprintf("OK %s %d\n", tmpDir, i)))
	})

	http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), mux)
}

func fetchImages(images []string, res *[][]byte) cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(urlstr),
		cdp.Sleep(2 * time.Second),
		cdp.WaitVisible(sel, cdp.ByID),
		cdp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, cdp.ByQuery),
		cdp.Screenshot(sel, res, cdp.NodeVisible, cdp.ByID),
	}
}
