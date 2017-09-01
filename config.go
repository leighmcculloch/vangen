package main

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"strings"
)

type config struct {
	Domain       string       `json:"domain"`
	Repositories []repository `json:"repositories"`
}

type repository struct {
	Prefix     string     `json:"prefix"`
	Subs       []string   `json:"subs"`
	Type       string     `json:"type"`
	URL        string     `json:"url"`
	SourceURLs sourceURLs `json:"source"`
	Website    website    `json:"website"`
}

type website struct {
	URL string `json:"url"`
}

func (r repository) Packages() []string {
	pkgs := []string{r.Prefix}
	for _, s := range r.Subs {
		pkgs = append(pkgs, path.Join(r.Prefix, s))
	}
	return pkgs
}

type sourceURLs struct {
	Home string `json:"home"`
	Dir  string `json:"dir"`
	File string `json:"file"`
}

func loadConfig(filename string) (config, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return config{}, err
	}

	var c config
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return config{}, err
	}

	for i := range c.Repositories {
		c.Repositories[i] = transformRepository(c.Repositories[i])
	}

	return c, nil
}

func transformRepository(r repository) repository {
	if strings.HasPrefix(r.URL, "https://github.com") {
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

	return r
}
