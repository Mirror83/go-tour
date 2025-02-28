package main

import (
	"image"
	"image/color"
)

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

type Image struct {
	Color                             color.Model
	TopLeftX, TopLeftY, Width, Height int
}

func (img Image) ColorModel() color.Model {
	return img.Color
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(img.TopLeftX, img.TopLeftY, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, 255, 255}
}
