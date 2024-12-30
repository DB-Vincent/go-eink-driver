package main

import (
	"fmt"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/graphics"
	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
	"github.com/DB-Vincent/go-eink-driver/pkg/utils"
)

func main() {
	spi, err := spi.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	display := display.New(spi, true)
	defer spi.Close()

	fmt.Println("Initializing display")
	display.Init()

	fmt.Println("Clearing display")
	display.Clear(utils.ColorWhite)

	fmt.Println("Drawing a rectangle")
	graphics.Rectangle(display.Canvas, 30, 62, 61, 125)
	display.DrawCanvas()

	display.Sleep()
}
