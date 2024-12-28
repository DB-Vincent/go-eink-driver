package utils

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// Text creates an image.Image containing the specified text
func Text(img *image.Gray, x, y int, label string) {
	point := fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: color.Gray{Y: 0}},
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	d.DrawString(label)
}
