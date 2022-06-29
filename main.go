package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var version = "<not set>"

func run() error {
	printHelp := flag.Bool("help", false, "print this help list")
	printVersion := flag.Bool("version", false, "print program version")
	verbose := flag.Bool("verbose", false, "print verbose output when run")
	filename := flag.String("config", "vangen.json", "vangen json configuration `filename`")
	outputDir := flag.String("out", "vangen/", "output `directory` that static files will be written to")
	noOverwrite := flag.Bool("no-overwrite", false, "If an output file already exists, stops with a non-zero return code")
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
		return nil
	}

	if *printVersion {
		fmt.Fprintln(os.Stderr, version)
		return nil
	}

	cf, err := os.Open(*filename)
	if err != nil {
		return err
	}

	c, err := parseConfig(cf)
	if err != nil {
		return err
	}

	err = os.MkdirAll(*outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("making dir %s: %w", *outputDir, err)
	}

	if c.Index {
		pathOut := filepath.Join(*outputDir, "index.html")
		f, err := os.Create(pathOut)
		if err != nil {
			return fmt.Errorf("writing file %s: %w", pathOut, err)
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
		err = generate_index(f, c.Domain, c.Repositories)
		if err != nil {
			return fmt.Errorf("generating index: %w", err)
		}

		err = f.Sync()
		if err != nil {
			return fmt.Errorf("flushing file %s: %w", pathOut, err)
		}
	}

	for _, r := range c.Repositories {
		for _, p := range r.Packages() {
			dirOut := filepath.Join(*outputDir, p)
			err = os.MkdirAll(dirOut, os.ModePerm)
			if err != nil {
				return fmt.Errorf("making dir %s: %v", dirOut, err)
			}

			pathOut := filepath.Join(dirOut, "index.html")
			if *noOverwrite {
				if _, err := os.Stat(pathOut); !os.IsNotExist(err) {
					if err == nil {
						return fmt.Errorf("cannot overwrite output file %s: %w", pathOut, err)
					}

					return fmt.Errorf("checking file %s: %w", pathOut, err)
				}
			}
			f, err := os.Create(pathOut)
			if err != nil {
				return fmt.Errorf("writing file %s: %v", pathOut, err)
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
			err = generate_package(f, c.Domain, c.DocsDomain, p, r)
			if err != nil {
				return fmt.Errorf("generating package %s: %w", p, err)
			}

			err = f.Sync()
			if err != nil {
				return fmt.Errorf("flushing file %s: %w", pathOut, err)
			}
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
