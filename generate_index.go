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
<title>{{$.Domain}} Go Modules</title>
<style>
* { font-family: sans-serif; }
body { margin-top: 0; }
.content { display: inline-block; }
</style>
</head>
<body>
<div class="content">

<h2>{{$.Domain}} Go Modules</h2>

<h3>Tools:</h3>

<ul>
{{range $_, $r := .MainRepositories -}}
<li>
<a href="/{{$r.Prefix}}">{{$r.Prefix}}</a>
{{if .Subs -}}<ul>{{end -}}
{{range $_, $s := .Subs -}}{{if not $s.Hidden -}}<li><a href="/{{$r.Prefix}}/{{$s.Name}}">{{$s.Name}}</a></li>{{end -}}{{end -}}
{{if .Subs -}}</ul>{{end -}}
</li>
{{end -}}
</ul>

<h3>Libraries:</h3>

<ul>
{{range $_, $r := .PackageRepositories -}}
<li>
<a href="/{{$r.Prefix}}">{{$r.Prefix}}</a>
{{if .Subs -}}<ul>{{end -}}
{{range $_, $s := .Subs -}}{{if not $s.Hidden -}}<li><a href="/{{$r.Prefix}}/{{$s.Name}}">{{$s.Name}}</a></li>{{end -}}{{end -}}
{{if .Subs -}}</ul>{{end -}}
</li>
{{end -}}
</ul>

<hr/>

Generated by <a href="https://4d63.com/vangen">vangen</a>.

</div>
</body>
</html>`

	tmpl, err := template.New("").Parse(html)
	if err != nil {
		return fmt.Errorf("error loading template: %v", err)
	}

	mainRepositories := []repository{}
	packageRepositories := []repository{}
	for _, r := range r {
		if r.Main {
			mainRepositories = append(mainRepositories, r)
		} else {
			packageRepositories = append(packageRepositories, r)
		}
	}

	data := struct {
		Domain              string
		MainRepositories    []repository
		PackageRepositories []repository
	}{
		Domain:              domain,
		MainRepositories:    mainRepositories,
		PackageRepositories: packageRepositories,
	}

	err = tmpl.ExecuteTemplate(w, "", data)
	if err != nil {
		return fmt.Errorf("generating template: %v", err)
	}

	return nil
}
