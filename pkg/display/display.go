package display

import (
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
)

type Display struct {
	Width  int
	Height int

	Spi *spi.SPI
}

// New creates a new display
func New(spi *spi.SPI) *Display {
	d := &Display{}

	d.Width = 122
	d.Height = 250

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
	d.Spi.SendCommand(0x4E) // SET_RAM_X_ADDRESS_COUNTER
	d.Spi.SendByte(byte(x & 0xFF))

	d.Spi.SendCommand(0x4F) // SET_RAM_Y_ADDRESS_COUNTER
	d.Spi.SendByte(byte(y & 0xFF))
	d.Spi.SendByte(byte((y >> 8) & 0xFF))
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

// TurnDisplayOn turns on normal display refresh
func (d *Display) TurnDisplayOn() {
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
	d.TurnDisplayOn()
}

// Sleep puts the display in sleep mode
func (d *Display) Sleep() {
	d.Spi.SendCommand(0x10) // Deep sleep mode
	d.Spi.SendByte(0x01)
	time.Sleep(2000 * time.Millisecond)
}

// DrawDot draws a dot on the screen at the given X & Y coordinate
func (d *Display) DrawDot(x, y int, color byte) {
	d.SetCursor(x, y)
	d.Spi.SendCommand(0x24) // Write RAM
	d.Spi.SendByte(color)
	d.TurnDisplayOn()
}
