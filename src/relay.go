package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin_number, err := strconv.Atoi(os.Getenv("RPIO_PIN"))
	if err != nil {
		pin_number = 21
	}

	interval, err := strconv.Atoi(os.Getenv("BLINK_INTERVAL"))
	if err != nil {
		interval = 1
	}

	pin := rpio.Pin(pin_number)
	pin.Output()

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		pin.Low()
		os.Exit(0)
	}()

	duration := time.Duration(interval) * time.Second

	for {
		pin.Toggle()
		if pin.Read() == 1 {
			fmt.Println("Blink is On")
		} else {
			fmt.Println("Blink is Off")
		}
		time.Sleep(duration)
	}
}

