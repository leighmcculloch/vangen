package main

import (
	"fmt"
	"html/template"
	"io"
)

func generate(w io.Writer, domain, pkg string, r repository) error {
	const html = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="{{.Domain}}/{{.Repository.Prefix}} {{.Repository.Type}} {{.Repository.URL}}">
<meta name="go-source" content="{{.Domain}}/{{.Repository.Prefix}} {{.Repository.SourceURLs.Home}} {{.Repository.SourceURLs.Dir}} {{.Repository.SourceURLs.File}}">
<meta http-equiv="refresh" content="0; url={{.HomeURL}}">
</head>
<body>
If you are not redirected, <a href="{{.HomeURL}}">click here</a>.
</body>`

	tmpl, err := template.New("").Parse(html)
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	var homeURL string
	if r.Website.URL != "" {
		homeURL = r.Website.URL
	} else {
		homeURL = fmt.Sprintf("https://godoc.org/%s/%s", domain, pkg)
	}

	data := struct {
		Domain     string
		Package    string
		Repository repository
		HomeURL    string
	}{
		Domain:     domain,
		Package:    pkg,
		Repository: r,
		HomeURL:    homeURL,
	}

	err = tmpl.ExecuteTemplate(w, "", data)
	if err != nil {
		return fmt.Errorf("generating template: %v", err)
	}

	return nil
}
