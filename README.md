# vangen

Static HTML generator for hosting Go repositories at a vanity import path.

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

1. Create a `vangen.json`.
   * [example-1.json](example-1.json) - Minimum required fields.
   * [example-2.json](example-2.json) - All fields.
2. Run `vangen`.
3. Host the files outputted in `vangen/` at your domain.
4. Try it out with `go get [domain]/[package]`
