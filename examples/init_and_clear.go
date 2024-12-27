package examples

import (
	"fmt"
	"time"

	"github.com/DB-Vincent/go-eink-driver/pkg/display"
	"github.com/DB-Vincent/go-eink-driver/pkg/spi"
)

func main() {
	spi, err := spi.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	display := display.New(spi)
	defer spi.Close()

	fmt.Println("Initializing display...")
	display.Init()

	fmt.Println("Clearing display...")
	display.Clear(0xFF)

	fmt.Println("Done. Sleeping in 5 seconds...")
	time.Sleep(5 * time.Second)
	display.Sleep()
}
