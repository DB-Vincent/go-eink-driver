package utils

import (
	"image"
	"image/color"
)

func Rectangle(x, y, width, height int) *image.RGBA {
	// Create image with correct bounds
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill rectangle
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img.Set(px, py, color.Black)
		}
	}

	return img
}
