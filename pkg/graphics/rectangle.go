package graphics

import (
	"image"
	"image/color"
	"image/draw"
)

func Rectangle(img *image.Gray, x, y, width, height int) {
	rectangle := image.Rect(x, y, x+width, y+height)
	draw.Draw(img, rectangle, &image.Uniform{C: color.Gray{Y: 0}}, image.Point{}, draw.Src)
}
