package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	help := flag.Bool("help", false, "")
	filename := flag.String("file", "vangen.json", "vangen json file")
	outputDir := flag.String("out", "vangen/", "output dir")
	flag.Parse()

	if *help {
		flag.Usage()
	}

	c, err := loadConfig(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	for _, cp := range c.Packages {
		dirOut := filepath.Join(*outputDir, cp.Name)
		pathOut := filepath.Join(dirOut, "index.html")

		out, err := htmlForPackage(c.Domain, cp.Name, cp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "generating html for package %s: %v", cp.Name, err)
			return
		}

		err = os.MkdirAll(dirOut, os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "making dir path %s: %v", dirOut, err)
			return
		}

		err = ioutil.WriteFile(pathOut, out, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "writing file %s: %v", pathOut, err)
			return
		}

		for _, sp := range cp.Subpackages {
			dirOut := filepath.Join(dirOut, sp)
			pathOut := filepath.Join(dirOut, "index.html")
			absName := cp.Name + "/" + sp

			out, err := htmlForPackage(c.Domain, absName, cp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "generating html for package %s: %v", absName, err)
				return
			}

			err = os.MkdirAll(dirOut, os.ModePerm)
			if err != nil {
				fmt.Fprintf(os.Stderr, "making dir path %s: %v", dirOut, err)
				return
			}

			err = ioutil.WriteFile(pathOut, out, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "writing file %s: %v", pathOut, err)
				return
			}
		}
	}
}

func htmlForPackage(domain, name string, cp configPackage) ([]byte, error) {
	const html = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="go-import" content="{{.Domain}}/{{.Package.Name}} {{.Package.VCS}} {{.Package.Repo}}">
<meta name="go-source" content="{{.Domain}}/{{.Package.Name}} {{.Package.HomeURL}} {{.Package.DirURL}} {{.Package.FileURL}}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/{{.Domain}}/{{.Name}}">
</head>
<body>
If you are not redirected, <a href="https://godoc.org/{{.Domain}}/{{.Name}}">click here</a>.
</body>`

	tmpl, err := template.New("").Parse(html)
	if err != nil {
		return nil, fmt.Errorf("error loading template: %v", err)
	}

	data := struct {
		Domain  string
		Name    string
		Package configPackage
	}{
		Domain:  domain,
		Name:    name,
		Package: cp,
	}

	out := bytes.Buffer{}
	err = tmpl.ExecuteTemplate(&out, "", data)
	if err != nil {
		return nil, fmt.Errorf("generating template: %v", err)
	}

	return out.Bytes(), nil
}
