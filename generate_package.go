package main

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

func generate_package(w io.Writer, domain, docs, pkg string, r repository) error {
	const html = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>{{.Domain}}/{{.Package}}</title>
<meta name="go-import" content="{{.Domain}}/{{.Repository.Prefix}} {{.Repository.Type}} {{.Repository.URL}}">
<meta name="go-source" content="{{.Domain}}/{{.Repository.Prefix}} {{.Repository.SourceURLs.Home}} {{.Repository.SourceURLs.Dir}} {{.Repository.SourceURLs.File}}">
<style>
* { font-family: sans-serif; }
body { margin-top: 0; }
.content { display: inline-block; }
code { display: block; font-family: monospace; font-size: 1em; background-color: #d5d5d5; padding: 1em; margin-bottom: 16px; }
ul { margin-top: 16px; margin-bottom: 16px; }
</style>
</head>
<body>
<div class="content">
<h2>{{.Domain}}/{{.Package}}</h2>
<code>go get {{.Domain}}/{{.Package}}</code>
<code>import "{{.Domain}}/{{.Package}}"</code>
Home: <a href="{{.HomeURL}}">{{.HomeURL}}</a><br/>
Source: <a href="{{.Repository.URL}}">{{.Repository.URL}}</a><br/>
{{if .Repository.Subs -}}Sub-packages:<ul>{{end -}}
{{range $_, $s := .Repository.Subs -}}{{if not $s.Hidden -}}<li><a href="/{{$.Repository.Prefix}}/{{$s.Name}}">{{$.Domain}}/{{$.Repository.Prefix}}/{{$s.Name}}</a></li>{{end -}}{{end -}}
{{if .Repository.Subs -}}</ul>{{end -}}
</div>
</body>
</html>`

	tmpl, err := template.New("").Parse(html)
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	var homeURL string
	if r.Website.URL != "" {
		homeURL = r.Website.URL
	} else {
		if docs == "" {
			docs = "pkg.go.dev"
		}
		homeURL = fmt.Sprintf("https://%s/%s/%s", docs, domain, pkg)
	}

	if strings.HasPrefix(r.URL, "https://github.com") || strings.HasPrefix(r.URL, "https://gitlab.com") {
		r.Type = "git"
		r.SourceURLs = sourceURLs{
			Home: r.URL,
			Dir:  r.URL + "/tree/master{/dir}",
			File: r.URL + "/blob/master{/dir}/{file}#L{line}",
		}
	}

	if r.SourceURLs.Home == "" {
		r.SourceURLs.Home = "_"
	}
	if r.SourceURLs.Dir == "" {
		r.SourceURLs.Dir = "_"
	}
	if r.SourceURLs.File == "" {
		r.SourceURLs.File = "_"
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
