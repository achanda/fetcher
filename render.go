package fetcher

import (
	"os"
	"text/template"
)

const templateText = `
<?xml version="1.0" encoding="UTF-8"?>
<urlset
	xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
		http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
{{range .}}
<url><loc>{{.Url}}</loc><lastmod>{{.TimeStamp}}</lastmod></url>
{{- end}}
</urlset>
`

func Render(urls []Url) {
	t := template.Must(template.New("sitemap").Parse(templateText))
	t.Execute(os.Stdout, urls)
}
