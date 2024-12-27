package spi

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

type SPI struct {
	Port spi.PortCloser
	Conn spi.Conn

	RstPin  gpio.PinIO
	DcPin   gpio.PinIO
	CsPin   gpio.PinIO
	BusyPin gpio.PinIO
}

const (
	spiSpeed    = physic.Hertz * 3000000
	spiMode     = spi.Mode0 // CPHL=0, CPOL=0
	bitsPerWord = 8
)

// New initializes the SPI connection
func New() (*SPI, error) {
	s := &SPI{}

	if _, err := host.Init(); err != nil {
		return nil, err
	}

	p, err := spireg.Open("SPI0.0")
	if err != nil {
		return nil, err
	}

	c, err := p.Connect(spiSpeed, spiMode, bitsPerWord)
	if err != nil {
		return nil, err
	}

	s.Port = p
	s.Conn = c

	s.RstPin = gpioreg.ByName("GPIO17")
	s.DcPin = gpioreg.ByName("GPIO25")
	s.CsPin = gpioreg.ByName("GPIO8")
	s.BusyPin = gpioreg.ByName("GPIO24")

	return s, nil
}

// Close closes the SPI connection
func (s *SPI) Close() error {
	return s.Port.Close()
}

// SendByte sets DC pin high and sends byte over SPI
func (s *SPI) SendByte(data byte) {
	s.DcPin.Out(true)
	s.CsPin.Out(false)
	c := []byte{data}
	r := make([]byte, len(c))
	s.Conn.Tx(c, r)
	s.CsPin.Out(true)
}

// SendBytes sends multiple bytes over SPI
func (s *SPI) SendBytes(data []byte) {
	s.DcPin.Out(true)
	s.CsPin.Out(false)
	r := make([]byte, len(data))
	s.Conn.Tx(data, r)
	s.CsPin.Out(true)
}

// SendCommand sets DC pin low and sends byte over SPI
func (s *SPI) SendCommand(command byte) {
	s.DcPin.Out(false)
	s.CsPin.Out(false)
	c := []byte{command}
	r := make([]byte, len(c))
	s.Conn.Tx(c, r)
	s.CsPin.Out(true)
}

// ReadBusy waits until the busy pin goes low
func (s *SPI) ReadBusy() {
	fmt.Println("Device busy")
	for s.BusyPin.Read() == gpio.High {
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("Device busy release")
}
