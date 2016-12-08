package fetcher

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

var fakeurls = []Url{
	{
		"https://google.com",
		time.Now().Format(time.RFC3339),
	},
	{
		"https://quora.com",
		time.Now().Format(time.RFC3339),
	},
}

var expected = `
<?xml version="1.0" encoding="UTF-8"?>
<urlset
	xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
		http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">

<url><loc>https://google.com</loc><lastmod>2016-12-08T23:20:41Z</lastmod></url>
<url><loc>https://quora.com</loc><lastmod>2016-12-08T23:20:41Z</lastmod></url>
</urlset>
`

func TestRender(t *testing.T) {
	b := new(bytes.Buffer)
	b.Reset()
	err := Render(b, fakeurls)
	if err != nil {
		t.Errorf("Did not expect an error")
	}

	gotLines := strings.Count(b.String(), "\n")
	expectedLines := strings.Count(expected, "\n")
	if gotLines != expectedLines {
		t.Errorf("Expected %v lines, got %v lines", expectedLines, gotLines)
	}
}
