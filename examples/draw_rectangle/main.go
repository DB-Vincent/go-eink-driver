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
	dot := utils.Rectangle(100, 40) // Create rectangle at (0,0) with width 100 and height 40
	display.DrawImage(dot, 50, 30)  // Draw rectangle at (50, 30)

	fmt.Println("Going back to sleep in 5 seconds...")
	time.Sleep(5 * time.Second)
	display.Sleep()
}
