package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var version = "<not set>"

func main() {
	printHelp := flag.Bool("help", false, "print this help list")
	printVersion := flag.Bool("version", false, "print program version")
	verbose := flag.Bool("verbose", false, "print verbose output when run")
	filename := flag.String("config", "vangen.json", "vangen json configuration `filename`")
	outputDir := flag.String("out", "vangen/", "output `directory` that static files will be written to")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Vangen is a tool for generating static HTML for hosting Go repositories at a vanity import path.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n")
		fmt.Fprintf(os.Stderr, "  vangen [-config=vangen.json] [-out=vangen/]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *printHelp {
		flag.Usage()
		return
	}

	if *printVersion {
		fmt.Fprintln(os.Stderr, version)
		return
	}

	cf, err := os.Open(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	c, err := parseConfig(cf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	for _, r := range c.Repositories {
		for _, p := range r.Packages() {
			dirOut := filepath.Join(*outputDir, p)
			err = os.MkdirAll(dirOut, os.ModePerm)
			if err != nil {
				fmt.Fprintf(os.Stderr, "making dir %s: %v", dirOut, err)
				return
			}

			pathOut := filepath.Join(dirOut, "index.html")
			f, err := os.Create(pathOut)
			if err != nil {
				fmt.Fprintf(os.Stderr, "writing file %s: %v", pathOut, err)
				return
			}
			defer func() {
				err := f.Close()
				if err != nil {
					fmt.Fprintf(os.Stderr, "closing file %s: %v", pathOut, err)
				}
			}()

			if *verbose {
				fmt.Fprintf(os.Stderr, "Writing %s\n", pathOut)
			}
			err = generate(f, c.Domain, p, r)
			if err != nil {
				fmt.Fprintf(os.Stderr, "generating package %s: %v", p, err)
				return
			}

			err = f.Sync()
			if err != nil {
				fmt.Fprintf(os.Stderr, "flushing file %s: %v", pathOut, err)
				return
			}
		}
	}
}
