package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	config "mechied.com/servocoolant/config"
	devicemanager "mechied.com/servocoolant/devicemanager"
	slog "mechied.com/servocoolant/logger"
)

type ServoCoolant struct {
	log           *log.Logger
	config        *config.Config
	deviceManager *devicemanager.DeviceManager
}

func main() {
	sc := ServoCoolant{}

	sc.log = slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Application ")
	sc.config = config.GetConfig(sc.log)

	sc.deviceManager = devicemanager.GetDeviceManager(sc.log, sc.config)

	sc.Run()

}

func (sc *ServoCoolant) Run() {
	sc.RegisterEndpoints()

	go sc.deviceManager.Servo1.TestServoWiggle()


	length, _ := sc.config.GetToolLength(12)
	sc.log.Info(fmt.Sprintf("getting tool length 12: %v", *length))

	sc.config.SetToolLength(12, 15.123)

	length, _ = sc.config.GetToolLength(12)
	sc.log.Info(fmt.Sprintf("getting tool length 12: %v", *length))

	sc.log.Fatal(http.ListenAndServe(":8080", sc.enableCors(sc.logRequest(http.DefaultServeMux))))
	//sc.deviceManager.RunAngleTest()
}
