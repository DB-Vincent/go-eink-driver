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

	point := fixed.Point26_6{X: fixed.I(0), Y: fixed.I(25)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(s)

	rotatedImg := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dy(), img.Bounds().Dx()))
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			rotatedImg.Set(img.Bounds().Dy()-1-y, x, img.At(x, y))
		}
	}
	img = rotatedImg

	return img
}
