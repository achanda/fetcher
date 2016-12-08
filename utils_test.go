package fetcher

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

type tokenTest struct {
	// The HTML to parse.
	html string
	// The string representations of the expected tokens, joined by '$'.
	expected string
}

var tokenTests = []tokenTest{
	{
		`<a href="http://www.w3schools.com/html/">Visit our HTML tutorial</a>`,
		"http://www.w3schools.com/html/",
	},
	{
		`<a href="http://www.w3schools.com/" target="_blank">Visit W3Schools!</a>`,
		"http://www.w3schools.com/",
	},
	{
		`<a href="">Visit W3Schools!</a>`,
		"",
	},
	{
		`<b blah></b>`,
		"",
	},
}

func TestExtractLink(t *testing.T) {
	for _, tt := range tokenTests {
		z := html.NewTokenizer(strings.NewReader(tt.html))
		z.Next()
		url := ExtractLink(z.Token())
		if url != tt.expected {
			t.Errorf("Expected %s, got %s", tt.expected, url)
		}
	}
}
