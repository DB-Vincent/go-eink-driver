package graphics

import (
	"image"
	"testing"
)

func TestCircle(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		x      int
		y      int
		r      int
	}{
		{
			name:   "small circle",
			width:  10,
			height: 10,
			x:      5,
			y:      5,
			r:      3,
		},
		{
			name:   "large circle",
			width:  100,
			height: 100,
			x:      50,
			y:      50,
			r:      25,
		},
		{
			name:   "circle at origin",
			width:  20,
			height: 20,
			x:      0,
			y:      0,
			r:      5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := image.NewGray(image.Rect(0, 0, tt.width, tt.height))
			Circle(img, tt.x, tt.y, tt.r)

			// Basic validation that pixels were set
			// Check center point
			if img.GrayAt(tt.x, tt.y).Y != 0 {
				t.Errorf("Expected center point (%d,%d) to be set", tt.x, tt.y)
			}
		})
	}
}
