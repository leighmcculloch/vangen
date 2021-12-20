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
	DocsDomain   string       `json:"docsDomain"`
	Index        bool         `json:"index"`
	Repositories []repository `json:"repositories"`
}

type repository struct {
	Prefix     string     `json:"prefix"`
	Subs       []sub      `json:"subs"`
	Type       string     `json:"type"`
	URL        string     `json:"url"`
	Main       bool       `json:"main"`
	Hidden     bool       `json:"hidden"`
	SourceURLs sourceURLs `json:"source"`
	Website    website    `json:"website"`
}

func (r repository) Packages() []string {
	pkgs := []string{r.Prefix}
	for _, s := range r.Subs {
		pkgs = append(pkgs, path.Join(r.Prefix, s.Name))
	}
	return pkgs
}

type sub struct {
	Name   string
	Hidden bool
}

func (s *sub) UnmarshalJSON(raw []byte) error {
	*s = sub{}

	err := json.Unmarshal(raw, &s.Name)
	if err == nil {
		return nil
	}

	subWithTags := struct {
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
	}{}
	err = json.Unmarshal(raw, &subWithTags)
	if err != nil {
		return err
	}
	*s = sub(subWithTags)
	return nil
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
