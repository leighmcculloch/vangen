package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Domain   string
	Packages []configPackage
}

type configPackage struct {
	Name        string `json:"name"`
	VCS         string `json:"vcs"`
	Repo        string
	HomeURL     string
	DirURL      string
	FileURL     string
	Subpackages []string
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

	for i := range c.Packages {
		if c.Packages[i].HomeURL == "" {
			c.Packages[i].HomeURL = "_"
		}
		if c.Packages[i].DirURL == "" {
			c.Packages[i].DirURL = "_"
		}
		if c.Packages[i].FileURL == "" {
			c.Packages[i].FileURL = "_"
		}
	}

	return c, nil
}
