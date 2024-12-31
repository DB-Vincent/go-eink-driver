package graphics

import (
	"image"
	"testing"
)

func TestImage(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		x, y    int
		wantErr bool
	}{
		{
			name:    "invalid path",
			path:    "nonexistent.png",
			x:       0,
			y:       0,
			wantErr: true,
		},
		{
			name:    "valid png",
			path:    "../../examples/draw_image/image.png",
			x:       0,
			y:       0,
			wantErr: false,
		},
		{
			name:    "with offset",
			path:    "../../examples/draw_image/image.png",
			x:       10,
			y:       20,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canvas := image.NewGray(image.Rect(0, 0, 100, 100))
			err := Image(canvas, tt.path, tt.x, tt.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("Image() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
