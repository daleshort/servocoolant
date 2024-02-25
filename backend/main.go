package main

import (
	"fmt"

	config "mechied.com/servocoolant/config"
	devicemanager "mechied.com/servocoolant/devicemanager"
	slog "mechied.com/servocoolant/logger"
	"net/http"
)

func main() {
	log := slog.GetLog()
	log.Info("Initializing Servo Coolant Nozzle Application")
	config := config.GetConfig(log)
	fmt.Println(config.GetVersion())
	devicemanager := devicemanager.GetDeviceManager(log, config)


	///home/dale/go/pkg/mod

	//devicemanager.RunRangeTest()
	devicemanager.RunAngleTest()

	http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}