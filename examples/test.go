package main

import (
	"fmt"
	"github.com/DB-Vincent/spi"
	"log"
)

func main() {
	// Create a new SPI object
	s := &spi.SPI{}

	// Initialize SPI connection
	err := s.New()
	if err != nil {
		log.Fatalf("Failed to initialize SPI: %v", err)
	}
	defer s.Close()

	// Test transmitting data
	writeData := []byte{0x01, 0x02, 0x03}    // Example data to write
	readData := make([]byte, len(writeData)) // Buffer to hold read data

	// Perform a transaction (write and read)
	err = s.Tx(writeData, readData)
	if err != nil {
		log.Fatalf("SPI transaction failed: %v", err)
	}

	// Output the received data (if any)
	fmt.Printf("Received Data: %v\n", readData)
}
