package graphics

import (
	"image"
	"image/draw"
	_ "image/png"
	"os"
)

func Image(canvas *image.Gray, path string, x, y int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	draw.Draw(canvas, canvas.Bounds(), img, image.Point{-x, -y}, draw.Over)
	return nil
}
