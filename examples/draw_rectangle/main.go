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
	rectangle := utils.Rectangle(50, 250)
	display.DrawImage(rectangle)

	fmt.Println("Going back to sleep in 5 seconds...")
	time.Sleep(5 * time.Second)
	display.Sleep()
}
