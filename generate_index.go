package main

import (
	"fmt"
	"html/template"
	"io"
)

func generate_index(w io.Writer, domain string, r []repository) error {
	const html = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<style>
* { font-family: sans-serif; }
</style>
</head>
<body>
<ul>
{{range $_, $r := .Repositories}}
<li><a href="//{{$.Domain}}/{{$r.Prefix}}">{{$.Domain}}/{{$r.Prefix}}</a></li>
{{range $_, $s := .Subs}}
<li><a href="//{{$.Domain}}/{{$r.Prefix}}/{{$s}}">{{$.Domain}}/{{$r.Prefix}}/{{$s}}</a></li>
{{end}}
{{end}}
</ul>
</body>
</html>`

	tmpl, err := template.New("").Parse(html)
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	data := struct {
		Domain       string
		Repositories []repository
	}{
		Domain:       domain,
		Repositories: r,
	}

	err = tmpl.ExecuteTemplate(w, "", data)
	if err != nil {
		return fmt.Errorf("generating template: %v", err)
	}

	return nil
}
