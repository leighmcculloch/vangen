package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseConfigNoIndex(t *testing.T) {
	r := strings.NewReader(`{}`)

	c, err := parseConfig(r)
	if err != nil {
		t.Fatal(err)
	}

	if g, w := c.Index, false; g != w {
		t.Errorf("Got index %#v, want %#v", g, w)
	}
}

func TestParseConfigIndex(t *testing.T) {
	r := strings.NewReader(`{
  "index": true
}`)

	c, err := parseConfig(r)
	if err != nil {
		t.Fatal(err)
	}

	if g, w := c.Index, true; g != w {
		t.Errorf("Got index %#v, want %#v", g, w)
	}
}

func TestParseConfigPackages(t *testing.T) {
	r := strings.NewReader(`{
  "repositories": [
    {
      "prefix": "foo",
      "subs": [
        "bar",
		"car",
		"car/dar",
		"ear/far"
      ]
    }
  ]
}`)

	c, err := parseConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	p := c.Repositories[0].Packages()

	e := []string{
		"foo",
		"foo/bar",
		"foo/car",
		"foo/car/dar",
		"foo/ear/far",
	}

	if !reflect.DeepEqual(p, e) {
		t.Errorf("Got packages %#v, want %#v", p, e)
	}
}

func TestParseConfigHiddenPackages(t *testing.T) {
	r := strings.NewReader(`{
  "repositories": [
    {
      "prefix": "foo",
      "subs": [
        "bar",
		"car",
		{ "name": "car/dar", "hidden": true },
		{ "name": "ear/far", "hidden": false }
      ]
    }
  ]
}`)

	c, err := parseConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	s := c.Repositories[0].Subs

	e := []sub{
		{Name: "bar"},
		{Name: "car"},
		{Name: "car/dar", Hidden: true},
		{Name: "ear/far", Hidden: false},
	}

	if !reflect.DeepEqual(s, e) {
		t.Errorf("Got packages %#v, want %#v", s, e)
	}
}

func TestParseConfigGithubMinimal(t *testing.T) {
	r := strings.NewReader(`{
  "domain": "4d63.com",
  "repositories": [
    {
      "prefix": "optional",
      "subs": [
        "template"
      ],
      "url": "https://github.com/leighmcculloch/go-optional"
    }
  ]
}`)

	e := config{
		Domain: "4d63.com",
		Repositories: []repository{
			{
				Prefix: "optional",
				Subs: []sub{
					{Name: "template"},
				},
				URL: "https://github.com/leighmcculloch/go-optional",
			},
		},
	}

	c, err := parseConfig(r)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(c, e) {
		t.Errorf("Got config %#v, want %#v", c, e)
	}
}

func TestParseConfigGithubComplete(t *testing.T) {
	r := strings.NewReader(`{
  "domain": "4d63.com",
  "repositories": [
    {
      "prefix": "optional",
      "subs": [
        "template"
      ],
      "type": "git",
      "url": "https://github.com/leighmcculloch/go-optional",
      "source": {
        "home": "https://github.com/leighmcculloch/go-optional",
        "dir": "https://github.com/leighmcculloch/go-optional/tree/master{/dir}",
        "file": "https://github.com/leighmcculloch/go-optional/blob/master{/dir}/{file}#L{line}"
      },
      "website": {
        "url": "https://github.com/leighmcculloch/go-optional"
      }
    }
  ]
}`)

	e := config{
		Domain: "4d63.com",
		Repositories: []repository{
			{
				Prefix: "optional",
				Subs: []sub{
					{Name: "template"},
				},
				Type: "git",
				URL:  "https://github.com/leighmcculloch/go-optional",
				SourceURLs: sourceURLs{
					Home: "https://github.com/leighmcculloch/go-optional",
					Dir:  "https://github.com/leighmcculloch/go-optional/tree/master{/dir}",
					File: "https://github.com/leighmcculloch/go-optional/blob/master{/dir}/{file}#L{line}",
				},
				Website: website{
					URL: "https://github.com/leighmcculloch/go-optional",
				},
			},
		},
	}

	c, err := parseConfig(r)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(c, e) {
		t.Errorf("Got config %#v, want %#v", c, e)
	}
}
