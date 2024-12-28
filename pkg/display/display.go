package display

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
)

type Display struct {
	Width  int
	Height int

	Canvas *image.Gray

	Spi *spi.SPI
}

// New creates a new display
func New(spi *spi.SPI) *Display {
	d := &Display{}

	d.Width = 122
	d.Height = 250

	d.Canvas = image.NewGray(image.Rect(0, 0, d.Width, d.Height))                                          // Create new canvas
	draw.Draw(d.Canvas, d.Canvas.Bounds(), &image.Uniform{C: color.Gray{Y: 255}}, image.Point{}, draw.Src) // Fill canvas with white

	d.Spi = spi

	return d
}

// Reset resets the display
func (d *Display) Reset() {
	d.Spi.RstPin.Out(true)
	time.Sleep(20 * time.Millisecond)
	d.Spi.RstPin.Out(false)
	time.Sleep(2 * time.Millisecond)
	d.Spi.RstPin.Out(true)
	time.Sleep(20 * time.Millisecond)
}

// SetWindow sets the display window
func (d *Display) SetWindow(xStart, yStart, xEnd, yEnd int) {
	d.Spi.SendCommand(0x44) // SET_RAM_X_ADDRESS_START_END_POSITION
	d.Spi.SendByte(byte((xStart >> 3) & 0xFF))
	d.Spi.SendByte(byte((xEnd >> 3) & 0xFF))

	d.Spi.SendCommand(0x45) // SET_RAM_Y_ADDRESS_START_END_POSITION
	d.Spi.SendByte(byte(yStart & 0xFF))
	d.Spi.SendByte(byte((yStart >> 8) & 0xFF))
	d.Spi.SendByte(byte(yEnd & 0xFF))
	d.Spi.SendByte(byte((yEnd >> 8) & 0xFF))
}

// SetCursor sets the cursor position
func (d *Display) SetCursor(x, y int) {
	d.Spi.SendCommand(0x4E) // Set RAM X address counter
	d.Spi.SendByte(byte(x / 8))
	d.Spi.SendCommand(0x4F) // Set RAM Y address counter
	d.Spi.SendByte(byte(y))
	d.Spi.ReadBusy() // Wait until the display is ready
}

// Init initializes the display
func (d *Display) Init() {
	d.Reset()
	d.Spi.ReadBusy()

	d.Spi.SendCommand(0x12) // SWRESET
	d.Spi.ReadBusy()

	d.Spi.SendCommand(0x01) // Driver output control
	d.Spi.SendByte(0xF9)
	d.Spi.SendByte(0x00)
	d.Spi.SendByte(0x00)

	d.Spi.SendCommand(0x11) // Datad.try mode
	d.Spi.SendByte(0x03)

	d.SetWindow(0, 0, d.Width-1, d.Height-1)
	d.SetCursor(0, 0)

	d.Spi.SendCommand(0x3C) // BorderWavefrom
	d.Spi.SendByte(0x05)

	d.Spi.SendCommand(0x21) // Display update control
	d.Spi.SendByte(0x00)
	d.Spi.SendByte(0x80)

	d.Spi.SendCommand(0x18)
	d.Spi.SendByte(0x80)

	d.Spi.ReadBusy()
}

// Refresh turns on normal display refresh
func (d *Display) Refresh() {
	d.Spi.SendCommand(0x22) // Display Update Control
	d.Spi.SendByte(0xF7)
	d.Spi.SendCommand(0x20) // Activate Display Update Sequence
	d.Spi.ReadBusy()
}

// Clear clears the display
func (d *Display) Clear(color byte) {
	lineWidth := d.Width / 8
	if d.Width%8 != 0 {
		lineWidth++
	}

	d.Spi.SendCommand(0x24)
	for i := 0; i < d.Height*lineWidth; i++ {
		d.Spi.SendByte(color)
	}
	d.Refresh()
}

// Sleep puts the display in sleep mode
func (d *Display) Sleep() {
	d.Spi.SendCommand(0x10) // Deep sleep mode
	d.Spi.SendByte(0x01)
	time.Sleep(2000 * time.Millisecond)
}

// DrawCanvas draws the canvas to the display
func (d *Display) DrawCanvas() {
	// Set cursor to start position
	d.SetCursor(0, 0)

	// Set RAM address bounds to full display
	d.SetWindow(0, 0, d.Width-1, d.Height-1)

	bytesPerRow := (d.Width + 7) / 8
	buffer := make([]byte, bytesPerRow*d.Height)

	for y := 0; y < d.Height; y++ {
		for x := 0; x < d.Width; x++ {
			if d.Canvas.GrayAt(x, y).Y > 128 {
				// Set the corresponding bit for black pixel
				buffer[y*bytesPerRow+x/8] |= 0x80 >> (x % 8)
			}
		}
	}

	// Write data
	d.Spi.SendCommand(0x24) // Write RAM
	d.Spi.SendBytes(buffer)

	d.Refresh()
}
