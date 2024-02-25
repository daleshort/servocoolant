package main

import (

	"net/http"
	config "mechied.com/servocoolant/config"
	devicemanager "mechied.com/servocoolant/devicemanager"
	slog "mechied.com/servocoolant/logger"
	log "github.com/sirupsen/logrus"
)

type ServoCoolant struct{
	log      *log.Logger
	config   *config.Config
	deviceManager *devicemanager.DeviceManager
}

func main() {
	sc := ServoCoolant{}

	sc.log= slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Applicatio ")
	sc.config = config.GetConfig(sc.log)

	sc.deviceManager = devicemanager.GetDeviceManager(sc.log, sc.config)

	sc.Run()

}



func (sc *ServoCoolant) Run (){
	sc.registerEndpoints()
	
    log.Fatal(http.ListenAndServe(":8080", nil))
	//sc.deviceManager.RunAngleTest()
}