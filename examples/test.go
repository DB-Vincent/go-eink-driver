package main

import (
	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
	"log"
)

func main() {
	s := &spi.SPI{}

	err := s.New()
	if err != nil {
		log.Fatalf("Failed to initialize SPI: %v", err)
	}
	defer s.Close()

	writeData := []byte{0x01}

	err = s.Tx(writeData, nil)
	if err != nil {
		log.Fatalf("SPI transaction failed: %v", err)
	}
}
