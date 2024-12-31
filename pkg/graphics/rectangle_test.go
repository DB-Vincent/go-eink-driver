package graphics

import (
	"image"
	"testing"
)

func TestRectangle(t *testing.T) {
	tests := []struct {
		name   string
		x      int
		y      int
		width  int
		height int
	}{
		{
			name:   "basic rectangle",
			x:      10,
			y:      10,
			width:  50,
			height: 30,
		},
		{
			name:   "zero position",
			x:      0,
			y:      0,
			width:  100,
			height: 100,
		},
		{
			name:   "small rectangle",
			x:      5,
			y:      5,
			width:  1,
			height: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := image.NewGray(image.Rect(0, 0, 200, 200))
			Rectangle(img, tt.x, tt.y, tt.width, tt.height)

			// Check corners are black
			if img.GrayAt(tt.x, tt.y).Y != 0 {
				t.Errorf("top-left corner should be black")
			}
			if img.GrayAt(tt.x+tt.width-1, tt.y+tt.height-1).Y != 0 {
				t.Errorf("bottom-right corner should be black")
			}
			if img.GrayAt(tt.x+tt.width-1, tt.y).Y != 0 {
				t.Errorf("top-right corner should be black")
			}
			if img.GrayAt(tt.x, tt.y+tt.height-1).Y != 0 {
				t.Errorf("bottom-left corner should be black")
			}
		})
	}
}
