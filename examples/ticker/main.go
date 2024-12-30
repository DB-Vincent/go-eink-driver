package main

import (
	"fmt"
	"time"

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

	var lastMinute, lastHour int

	for {
		currentTime := time.Now()
		currentMinute := currentTime.Minute()
		currentHour := currentTime.Hour()

		if currentMinute != lastMinute || currentHour != lastHour {
			display.Clear(utils.ColorWhite)

			timeString := currentTime.Format("15:04")
			graphics.Text(display.Canvas, 0, 0, timeString)

			display.DrawCanvas()
			display.ClearCanvas()

			lastMinute = currentMinute
			lastHour = currentHour
		}

		time.Sleep(1 * time.Second)
	}
}
