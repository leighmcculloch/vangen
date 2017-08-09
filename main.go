package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	help := flag.Bool("help", false, "")
	filename := flag.String("file", "vangen.json", "vangen json file")
	outputDir := flag.String("out", "vangen/", "output dir")
	flag.Parse()

	if *help {
		flag.Usage()
	}

	c, err := loadConfig(*filename)
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
