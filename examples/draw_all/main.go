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
	if err = graphics.Image(display.Canvas, "./image.png"); err != nil {
		fmt.Printf("%v", err)
	}
	graphics.Text(display.Canvas, 0, 0, "Hello, world!")
	graphics.Rectangle(display.Canvas, 0, 50, 50, 50)
	graphics.Circle(display.Canvas, 25, 150, 25)

	display.DrawCanvas()
	display.Sleep()
}
