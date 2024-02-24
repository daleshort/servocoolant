package main

import (
	"fmt"

	config "mechied.com/servocoolant/config"
	slog "mechied.com/servocoolant/logger"
)

func main() {
	log := slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Application  ")
	config := config.GetConfig(log)
	fmt.Println(config.GetVersion())


}
