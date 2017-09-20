# vangen
[![Linux/OSX Build Status](https://img.shields.io/travis/leighmcculloch/vangen.svg?label=linux%20%26%20osx)](https://travis-ci.org/leighmcculloch/vangen)
[![Windows Build Status](https://img.shields.io/appveyor/ci/leighmcculloch/vangen.svg?label=windows)](https://ci.appveyor.com/project/leighmcculloch/vangen)
[![Codecov](https://img.shields.io/codecov/c/github/leighmcculloch/vangen.svg)](https://codecov.io/gh/leighmcculloch/vangen)
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/vangen)](https://goreportcard.com/report/github.com/leighmcculloch/vangen)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/leighmcculloch/vangen)

Vangen is a tool for generating static HTML for Go vanity import paths.

Go vanity import paths work by serving a HTML file that tells the `go get` tool where to download the source from. You can still host the source code at Github, BitBucket, but the vanity URL gives you portability and other benefits.

## Why
* Maintain Go vanity import paths with a simple definition file `vangen.json`.
* Host Go vanity import paths using static hosting. No need for Google AppEngine, Heroku, etc. Host the files on Github Pages, AWS S3, Google Cloud Storage, etc.

## Install

### Source

```
go get 4d63.com/vangen
```

### Mac

```
curl -o /usr/local/bin/vangen https://raw.githubusercontent.com/leighmcculloch/vangen/binaries/mac/amd64/vangen && chmod +x /usr/local/bin/vangen
```

### Linux

```
curl -o /usr/local/bin/vangen https://raw.githubusercontent.com/leighmcculloch/vangen/binaries/linux/amd64/vangen && chmod +x /usr/local/bin/vangen
```

### Windows

[Download the executable](https://raw.githubusercontent.com/leighmcculloch/vangen/binaries/windows/amd64/vangen.exe), and save it to your path.

## Usage

1. Create a `vangen.json` (see examples below)
2. Run `vangen`
3. Host the files outputted in `vangen/` at your domain
4. Try it out with `go get [domain]/[package]`

```
$ vangen -help
Vangen is a tool for generating static HTML for Go vanity import paths.

Usage:

  vangen [-file=vangen.json] [-out=vangen/]

Flags:

  -file filename
        vangen json configuration filename (default "vangen.json")
  -help
        print this help list
  -out directory
        output directory that static files will be written to (default "vangen/")
  -version
        print program version
```

## Examples

### Minimal

```json
{
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
}
```

### All fields

```json
{
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
        "url": "https://github.com/leighmcculoch/go-optional"
      }
    }
  ]
}
```

