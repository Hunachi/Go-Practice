// Implemented by Hunachi

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	x int
	y int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0,0,img.x,img.y)
}

func (img Image) At(x, y int) color.Color {
	v := x + y
	return color.RGBA{uint8(255-v), uint8(v), 255, 255}
}

func main() {
	m := Image{100,100}
	pic.ShowImage(m)
}