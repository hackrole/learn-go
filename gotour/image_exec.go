package main

import (
	"golang.org/x/tour/pic"
)

type Image struct{}

func (img *Image) Bounds()

func main() {
	m := Image{}
	pic.ShowImage(m)
}
