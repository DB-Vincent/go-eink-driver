package display

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
)

const (
	physicalWidth  = 122
	physicalHeight = 250
)

type Display struct {
	Width  int
	Height int

	Canvas      *image.Gray
	IsLandscape bool

	Spi *spi.SPI
}

// New creates a new display
func New(spi *spi.SPI, landscape bool) *Display {
	d := &Display{}
	d.IsLandscape = landscape

	// In landscape mode, swap width and height for the canvas
	if landscape {
		d.Width = physicalHeight
		d.Height = physicalWidth
	} else {
		d.Width = physicalWidth
		d.Height = physicalHeight
	}

	d.ClearCanvas()

	d.Spi = spi
	return d
}

func (d *Display) ClearCanvas() {
	d.Canvas = image.NewGray(image.Rect(0, 0, d.Width, d.Height))
	draw.Draw(d.Canvas, d.Canvas.Bounds(), &image.Uniform{C: color.Gray{Y: 255}}, image.Point{}, draw.Src)
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
	d.Spi.SendCommand(0x11) // Data entry mode
	d.Spi.SendByte(0x03)
	d.Spi.SendCommand(0x44) // SET_RAM_X_ADDRESS_START_END_POSITION
	d.Spi.SendByte(0x00)
	d.Spi.SendByte(0x0F)
	d.Spi.SendCommand(0x45) // SET_RAM_Y_ADDRESS_START_END_POSITION
	d.Spi.SendByte(0x00)
	d.Spi.SendByte(0x00)
	d.Spi.SendByte(0xF9)
	d.Spi.SendByte(0x00)
	d.Spi.SendCommand(0x3C) // BorderWavefrom
	d.Spi.SendByte(0x05)
	d.Spi.SendCommand(0x21) // Display update control
	d.Spi.SendByte(0x00)
	d.Spi.SendByte(0x80)
	d.Spi.SendCommand(0x18)
	d.Spi.SendByte(0x80)
	d.Spi.ReadBusy()
}

// DrawCanvas draws the canvas to the display
func (d *Display) DrawCanvas() {
	bytesPerRow := (physicalWidth + 7) / 8
	buffer := make([]byte, bytesPerRow*physicalHeight)

	// Fill buffer based on orientation
	if d.IsLandscape {
		// Rotate 90 degrees clockwise when filling buffer
		for y := 0; y < physicalHeight; y++ {
			for x := 0; x < physicalWidth; x++ {
				canvasX := physicalHeight - 1 - y
				canvasY := x

				if d.Canvas.GrayAt(canvasX, canvasY).Y > 128 {
					buffer[y*bytesPerRow+x/8] |= 0x80 >> (x % 8)
				}
			}
		}
	} else {
		// Normal orientation
		for y := 0; y < physicalHeight; y++ {
			for x := 0; x < physicalWidth; x++ {
				if d.Canvas.GrayAt(x, y).Y > 128 {
					buffer[y*bytesPerRow+x/8] |= 0x80 >> (x % 8)
				}
			}
		}
	}

	// Write to display
	d.Spi.SendCommand(0x24) // Write RAM
	d.Spi.SendBytes(buffer)
	d.Refresh()
}

// Refresh turns on display update
func (d *Display) Refresh() {
	d.Spi.SendCommand(0x22)
	d.Spi.SendByte(0xF7)
	d.Spi.SendCommand(0x20)
	d.Spi.ReadBusy()
}

// Clear clears the display
func (d *Display) Clear(color byte) {
	lineWidth := physicalWidth / 8
	if physicalWidth%8 != 0 {
		lineWidth++
	}
	d.Spi.SendCommand(0x24)
	for i := 0; i < physicalHeight*lineWidth; i++ {
		d.Spi.SendByte(color)
	}
	d.Refresh()
}

// Sleep puts the display in sleep mode
func (d *Display) Sleep() {
	d.Spi.SendCommand(0x10)
	d.Spi.SendByte(0x01)
	time.Sleep(2000 * time.Millisecond)
}
