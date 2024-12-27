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
	img := image.NewRGBA(image.Rect(0, 0, 122, 255))

	point := fixed.Point26_6{X: 0, Y: 0}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)

	return img
}
