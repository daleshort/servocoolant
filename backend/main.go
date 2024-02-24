package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
	config "mechied.com/servocoolant/config"
	slog "mechied.com/servocoolant/logger"
)

func main() {
	log := slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Application ")
	config := config.GetConfig(log)
	fmt.Println(config.GetVersion())

	err := rpio.Open()

	if err != nil {
		log.Error(fmt.Sprintf("error opening pin %v", err))
	}

	pin :=rpio.Pin(13)
	pin.Input()
	pin.PullDown()

	for {

		res := pin.Read()

		log.Debug(fmt.Sprintf("pin is %v", res))
		time.Sleep(time.Millisecond*1000)
	}
	defer rpio.Close()
}
