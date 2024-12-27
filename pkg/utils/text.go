package utils

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func Text(label string) *image.RGBA {
	// Create image with correct bounds
	img := image.NewRGBA(image.Rect(0, 0, 200, 50)) // Adjust bounds as needed

	// Set starting point for text
	point := fixed.Point26_6{X: fixed.I(10), Y: fixed.I(20)} // Adjust coordinates as needed

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)

	return img
}
