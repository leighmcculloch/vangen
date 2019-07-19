package main

import (
	"encoding/json"
	"errors"
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
	Subs       []sub      `json:"subs"`
	Type       string     `json:"type"`
	URL        string     `json:"url"`
	Main       bool       `json:"main"`
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
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
}

func (s *sub) UnmarshalJSON(raw []byte) error {
	var v interface{}

	err := json.Unmarshal(raw, &v)
	if err != nil {
		return err
	}

	switch t := v.(type) {
	case string:
		s.Name = t

	case map[string]interface{}:
		s.Name = t["name"].(string)
		s.Hidden = t["hidden"].(bool)

	default:
		return errors.New("cannot unmarshal object into Go struct of type sub")
	}

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
