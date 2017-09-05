package main

import (
	"reflect"
	"strings"
	"testing"
)

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

func TestParseConfigGitubMinimal(t *testing.T) {
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
			repository{
				Prefix: "optional",
				Subs: []string{
					"template",
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
			repository{
				Prefix: "optional",
				Subs: []string{
					"template",
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
