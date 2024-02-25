package main

import (
	"fmt"

	config "mechied.com/servocoolant/config"
	devicemanager "mechied.com/servocoolant/devicemanager"
	slog "mechied.com/servocoolant/logger"
)

func main() {
	log := slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Application")
	config := config.GetConfig(log)
	fmt.Println(config.GetVersion())
	devicemanager := devicemanager.GetDeviceManager(log, config)


	///home/dale/go/pkg/mod

	devicemanager.RunRangeTest()
	devicemanager.RunAngleTest()
}
