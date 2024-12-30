package graphics

import (
	"image"
	"image/color"
)

// Circle creates an image.Image containing a circle
func Circle(img *image.Gray, x, y, r int) {
	x0, y0, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)
	c := color.Gray{Y: 0}

	for x0 >= y0 {
		for i := x - x0; i <= x+x0; i++ {
			img.Set(i, y+y0, c)
			img.Set(i, y-y0, c)
		}
		for i := x - y0; i <= x+y0; i++ {
			img.Set(i, y+x0, c)
			img.Set(i, y-x0, c)
		}

		if err <= 0 {
			y0++
			err += dy
			dy += 2
		}
		if err > 0 {
			x0--
			dx += 2
			err += dx - (r * 2)
		}
	}
}
