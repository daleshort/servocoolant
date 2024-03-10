package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	automanager "mechied.com/servocoolant/automanager"
	config "mechied.com/servocoolant/config"
	devicemanager "mechied.com/servocoolant/devicemanager"
	slog "mechied.com/servocoolant/logger"
)

type ServoCoolant struct {
	log           *log.Logger
	config        *config.Config
	deviceManager *devicemanager.DeviceManager
	autoManager   *automanager.AutoManager
}

func main() {
	sc := ServoCoolant{}

	sc.log = slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Application ")
	sc.config = config.GetConfig(sc.log)

	sc.deviceManager = devicemanager.GetDeviceManager(sc.log, sc.config)
	sc.autoManager = automanager.GetAutoManager(sc.log, sc.config, sc.deviceManager)
	sc.Run()

}

func (sc *ServoCoolant) Run() {
	sc.RegisterEndpoints()

	//go sc.deviceManager.Servo1.TestServoWiggle()

	// length, _ := sc.config.GetToolLength(12)
	// sc.log.Info(fmt.Sprintf("getting tool length 12: %v", *length))

	sc.config.SetToolLength(12, 2.123)
	sc.autoManager.CalculateAngleForToolLength(2)
	sc.config.SetToolLength(14, 2.123)
	sc.autoManager.CalculateAngleForToolLength(12)


	// length, _ = sc.config.GetToolLength(12)
	// sc.log.Info(fmt.Sprintf("getting tool length 12: %v", *length))

	sc.log.Fatal(http.ListenAndServe(":80", sc.enableCors(sc.logRequest(http.DefaultServeMux))))
	//sc.deviceManager.RunAngleTest()
}
