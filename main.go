package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// init function register the format to open/read image files and perform
// operation on the files.
func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

// main function
func main() {
	imgfile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("usage: %s <imagefile>\n", os.Args[0])
		os.Exit(-1)
	}
	defer imgfile.Close()

	img, _, err := image.Decode(imgfile)

	// Compute SIFT descriptor
	Sift(img)
}
