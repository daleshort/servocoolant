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

	log.Fatal(http.ListenAndServe(":8080", nil))
	//sc.deviceManager.RunAngleTest()
}

func (sc *ServoCoolant) RegisterEndpoints() {
	http.HandleFunc("/", sc.handler)

}

func (sc *ServoCoolant) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
