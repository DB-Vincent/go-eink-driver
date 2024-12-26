package spi

import (
	"errors"

	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

type SPI struct {
	port spi.PortCloser
	conn spi.Conn
}

const (
	spiSpeed    = physic.MegaHertz * 1
	spiMode     = spi.Mode3
	bitsPerWord = 8
)

func (s *SPI) New() error {
	if _, err := host.Init(); err != nil {
		return err
	}

	p, err := spireg.Open("SPI0.0")
	if err != nil {
		return err
	}

	c, err := p.Connect(spiSpeed, spiMode, bitsPerWord)
	if err != nil {
		return err
	}

	s.port = p
	s.conn = c

	return nil
}

func (s *SPI) Tx(write, read []byte) error {
	if s.conn == nil {
		return errors.New("SPI connection not initialized")
	}
	return s.conn.Tx(write, read)
}

func (s *SPI) Close() error {
	return s.port.Close()
}
