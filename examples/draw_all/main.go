package main

import (
	"fmt"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
	"github.com/DB-Vincent/go-eink-driver/pkg/utils"
)

func main() {
	spi, err := spi.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	display := display.New(spi)
	defer spi.Close()

	fmt.Println("Initializing display")
	display.Init()

	fmt.Println("Clearing display")
	display.Clear(utils.ColorWhite)

	fmt.Println("Drawing a rectangle")
	utils.Text(display.Canvas, 0, 0, "Hello, world!")
	utils.Rectangle(display.Canvas, 0, 50, 50, 50)
	utils.Circle(display.Canvas, 25, 150, 25)

	display.DrawCanvas()
	display.Sleep()
}
