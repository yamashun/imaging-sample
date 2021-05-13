package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

const (
	MAX_WIDTH = 300
)

func resize(path string) {
	srcImage, err := imaging.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	if srcImage.Bounds().Dx() > MAX_WIDTH {
		resizedImage := imaging.Resize(srcImage, MAX_WIDTH, 0, imaging.Lanczos)

		err = imaging.Save(resizedImage, path)
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("ERROR: 引数を指定してください。")
	}

	dirPath := os.Args[1]
	files, _ := ioutil.ReadDir(dirPath)

	for _, fileInfo := range files {
		resize(dirPath + fileInfo.Name())
	}
}
