package utils

import (
	"image"
	"image/color"
)

// Circle creates an image.Image containing a circle
func Circle(img *image.Gray, x0, y0, r int) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)
	c := color.Gray{Y: 0}

	for x >= y {
		for i := x0 - x; i <= x0+x; i++ {
			img.Set(i, y0+y, c)
			img.Set(i, y0-y, c)
		}
		for i := x0 - y; i <= x0+y; i++ {
			img.Set(i, y0+x, c)
			img.Set(i, y0-x, c)
		}

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}
