package main

import (
	"fmt"
	"time"

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
	display.DrawImage(rect, 125, 0)
	display.DrawImage(rect, 85, 15)
	display.DrawImage(rect, 165, 15)

	display.Refresh()

	fmt.Println("Going back to sleep in 5 seconds...")
	time.Sleep(5 * time.Second)
	display.Sleep()
}
