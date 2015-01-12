package main

import (
	"fmt"
	"image"
)

const (
	S      = 3
	SIGMA0 = 1.6
	MINW   = 20
)

//
func Sift(w, h int, img image.Image) {
	// calculate the number of octaves
	var width int
	if w < h {
		width = w
	} else {
		width = h
	}
	fmt.Printf("Width: %d, Height: %d\n", w, h)

	var numOct int
	for numOct = 0; width > MINW; width /= 2 {
		numOct++
	}
	fmt.Println("The number of octaves: ", numOct)

	getKeypoints(img)
	localizeKeypoints(img)
}

//
func getKeypoints(img image.Image) {
	fmt.Println("step1: Keypoint detection")
}

//
func localizeKeypoints(img image.Image) {
	fmt.Println("step2: Localize")
}
