package utils

import (
	"image"
	"image/color"
	"image/draw"
)

func Rectangle(x1, y1, x2, y2 int) *image.Gray {
	// Create a black-and-white image (Grayscale)
	width, height := 250, 122
	img := image.NewGray(image.Rect(0, 0, width, height))

	// Define the rectangle color (black in this case)
	black := color.Gray{Y: 0}   // Y=0 is black
	white := color.Gray{Y: 255} // Y=255 is white

	// Fill the entire image with white (background color)
	draw.Draw(img, img.Bounds(), &image.Uniform{C: white}, image.Point{}, draw.Src)

	// Define the rectangle's position and size
	rect := image.Rect(x1, y1, x2, y2)

	// Draw a black rectangle within the image
	draw.Draw(img, rect, &image.Uniform{C: black}, image.Point{}, draw.Over)

	return img
}
