package main

import (
	"fmt"
	"image"
)

//
func Sift(image image.Image) {
	getKeypoints(image)
	localizeKeypoints(image)
}

//
func getKeypoints(image image.Image) {
	fmt.Println("step1: Keypoint detection")
}

//
func localizeKeypoints(image image.Image) {
	fmt.Println("step2: Localize")
}
