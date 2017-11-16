package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"path"
	"sort"
)

type config struct {
	Domain       string       `json:"domain"`
	Index        bool         `json:"index"`
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

type website struct {
	URL string `json:"url"`
}

func parseConfig(r io.Reader) (config, error) {
	bytes, err := ioutil.ReadAll(r)

	if err != nil {
		return config{}, err
	}

	var c config
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return config{}, err
	}

	sort.Slice(c.Repositories, func(i, j int) bool {
		return c.Repositories[i].Prefix < c.Repositories[j].Prefix
	})

	return c, nil
}
