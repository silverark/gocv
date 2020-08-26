// What it does:
//
// 	This program outputs the current OpenVINO IE library version to the console.
//
// How to run:
//
// 		go run ./cmd/openvino/ie/version/main.go
//
// +build example

package main

import (
	"fmt"

	"github.com/silverark/gocv"
	"github.com/silverark/gocv/openvino/ie"
)

func main() {
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("OpenVINO Inference Engine version: %s\n", ie.Version())
}
