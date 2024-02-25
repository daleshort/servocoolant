package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (sc *ServoCoolant) Run() {
	sc.RegisterEndpoints()

	log.Fatal(http.ListenAndServe(":8080", nil))
	//sc.deviceManager.RunAngleTest()
}

func (sc *ServoCoolant) RegisterEndpoints() {
	http.HandleFunc("/", sc.handler)

}

func (sc *ServoCoolant) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I lo ve %s!", r.URL.Path[1:])
}
