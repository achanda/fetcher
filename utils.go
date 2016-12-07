package fetcher

import (
	"golang.org/x/net/html"
)

func ExtractLink(t html.Token) (href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}
	return href
}
