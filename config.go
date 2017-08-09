package main

import (
	"encoding/json"
	"io/ioutil"
	"path"
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
		if c.Repositories[i].SourceURLs.Home == "" {
			c.Repositories[i].SourceURLs.Home = "_"
		}
		if c.Repositories[i].SourceURLs.Dir == "" {
			c.Repositories[i].SourceURLs.Dir = "_"
		}
		if c.Repositories[i].SourceURLs.File == "" {
			c.Repositories[i].SourceURLs.File = "_"
		}
	}

	return c, nil
}
