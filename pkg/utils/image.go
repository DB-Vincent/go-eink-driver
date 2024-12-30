package utils

import (
	"image"
	"image/draw"
	_ "image/png"
	"os"
)

func Image(canvas *image.Gray, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	draw.Draw(canvas, canvas.Bounds(), img, image.Point{}, draw.Over)
	return nil
}
