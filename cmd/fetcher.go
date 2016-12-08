package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/achanda/fetcher"
	"golang.org/x/net/html"
)

// Extract all links from a given page as string
func crawl(uri string, ch chan string, chFinished chan bool) {
	givenUrl, err := url.Parse(uri)
	if err != nil {
		fmt.Println("Given URL does not look valid " + uri)
		return
	}
	resp, err := http.Get(uri)

	defer func() {
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("Failed to crawl " + uri)
		return
	}

	b := resp.Body
	defer b.Close()

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			eurl := fetcher.ExtractLink(t)
			if eurl == "" {
				continue
			}
			extractedUrl, _ := url.Parse(eurl)

			// Ignore external or relative links
			if extractedUrl.IsAbs() && extractedUrl.Host == givenUrl.Host {
				ch <- eurl
			}
		}
	}
}

func main() {
	seenUrls := make(map[string]bool)
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("usage: fetcher url")
		os.Exit(2)
	}

	// Channels
	chUrls := make(chan string)
	chFinished := make(chan bool)

	defer close(chUrls)
	defer close(chFinished)

	// Kick off the crawl process (concurrently)
	go crawl(args[0], chUrls, chFinished)

	// Subscribe to both channels
	for {
		select {
		case url := <-chUrls:
			seenUrls[url] = true
		case <-chFinished:
			var urls []fetcher.Url
			for url, _ := range seenUrls {
				urls = append(urls, fetcher.Url{Url: url, TimeStamp: time.Now().Format(time.RFC3339)})
			}
			fetcher.Render(urls)
			os.Exit(0)
		}
	}
}
