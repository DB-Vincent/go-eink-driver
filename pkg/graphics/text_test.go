package graphics

import (
	"image"
	"testing"
)

func TestText(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		text     string
		wantSize image.Rectangle
	}{
		{
			name: "basic text",
			x:    10,
			y:    20,
			text: "Test",
			wantSize: image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: 50, Y: 50},
			},
		},
		{
			name: "empty text",
			x:    0,
			y:    0,
			text: "",
			wantSize: image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: 50, Y: 50},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := image.NewGray(tt.wantSize)
			Text(img, tt.x, tt.y, tt.text)

			if img.Bounds() != tt.wantSize {
				t.Errorf("Text() resulted in image size = %v, want %v", img.Bounds(), tt.wantSize)
			}
		})
	}
}
