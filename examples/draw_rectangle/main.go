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
	rect := utils.Rectangle(50, 50)
	display.DrawImage(rect, 16, 26)
	display.DrawImage(rect, 36, 102)
	display.DrawImage(rect, 56, 178)

	display.Refresh()
	display.Sleep()
}
