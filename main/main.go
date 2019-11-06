package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is for demo
type Image struct{}

// ColorModel is for demo
func (im *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds is for demo
func (im *Image) Bounds() image.Rectangle {
	w := 256
	h := 256
	return image.Rect(0, 0, w, h)
}

// At is for demo
func (im *Image) At(x, y int) color.Color {
	// v := (x + y) / 2
	v := x * y
	// v := x ^ y

	return color.RGBA{uint8(v), uint8(v), 255, 255}
}

func main() {
	m := &Image{}
	pic.ShowImage(m)
}
