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
<meta http-equiv="refresh" content="0; url=https://godoc.org/{{.Domain}}/{{.Package}}">
</head>
<body>
If you are not redirected, <a href="https://godoc.org/{{.Domain}}/{{.Package}}">click here</a>.
</body>`

	tmpl, err := template.New("").Parse(html)
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	data := struct {
		Domain     string
		Package    string
		Repository repository
	}{
		Domain:     domain,
		Package:    pkg,
		Repository: r,
	}

	err = tmpl.ExecuteTemplate(w, "", data)
	if err != nil {
		return fmt.Errorf("generating template: %v", err)
	}

	return nil
}
