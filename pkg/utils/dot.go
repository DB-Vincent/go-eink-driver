package utils

import (
	"image"
	"image/color"
)

// Dot creates an image with a single dot at the specified coordinates.
func Dot(x, y int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 300, 400))
	img.Set(x, y, color.Black)

	return img
}
