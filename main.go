package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/disintegration/imaging"
)

const (
	MAX_WIDTH = 300
	OUT_PATH  = "tmp/"
)

func resize(path string) {
	srcImage, err := imaging.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	width := srcImage.Bounds().Dx()
	if width > MAX_WIDTH {
		resizedImage := imaging.Resize(srcImage, MAX_WIDTH, 0, imaging.Lanczos)

		err = imaging.Save(resizedImage, OUT_PATH+filepath.Base(path))
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	}
}

func main() {
	flag.Parse()
	paths := flag.Args()
	for _, path := range paths {
		resize(path)
	}
}
