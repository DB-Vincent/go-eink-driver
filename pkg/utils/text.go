package utils

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// Text creates an image.Image containing the specified text
func Text(s string) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 7*len(s)+55, 25))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	point := fixed.Point26_6{X: fixed.I(0), Y: fixed.I(10)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(s)

	return img
}
