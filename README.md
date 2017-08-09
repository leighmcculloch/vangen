# vangen

Static HTML generator for hosting Go repositories at a vanity URL.

Go vanity URLs work by serving a HTML file that tells the `go get` tool where to download the source from. You can still host the source code at Github, BitBucket, but the vanity URL gives you portability and other benefits.

## Why
* Maintain Go vanity URLs with a simple definition file `vangen.json`.
* Host Go vanity URLs using static hosting. No need for Google AppEngine, Heroku, etc. Host the files on Github Pages, AWS S3, Google Cloud Storage, etc.

## Install

`go get 4d63.com/vangen`

## Usage

1. Create a `vangen.json`.
   * [example-1.json](example-1.json) - Minimum required fields.
   * [example-2.json](example-2.json) - All fields.
2. Run `vangen`.
3. Host the files outputted in `vangen/` at your domain.
4. Try it out with `go get [domain]/[package]`
