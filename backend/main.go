package main

import (
	log "github.com/sirupsen/logrus"
	config "mechied.com/servocoolant/config"
	devicemanager "mechied.com/servocoolant/devicemanager"
	slog "mechied.com/servocoolant/logger"
)

func main() {
	sc := ServoCoolant{}

	sc.log = slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Application ")
	sc.config = config.GetConfig(sc.log)

	sc.deviceManager = devicemanager.GetDeviceManager(sc.log, sc.config)

	sc.Run()

}
