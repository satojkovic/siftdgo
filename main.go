package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
)

// init function register the format to open/read image files and perform
// operation on the files.
func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

// main function
func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Fprintf(
			os.Stderr,
			"usage: %s <imagefile>\n",
			os.Args[0],
		)
		os.Exit(-1)
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	conf, _, err := image.DecodeConfig(bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}

	// Compute SIFT descriptor
	Sift(conf.Width, conf.Height, img)
}
