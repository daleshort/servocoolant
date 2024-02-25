package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (sc *ServoCoolant) Run() {
	sc.RegisterEndpoints()

	log.Fatal(http.ListenAndServe(":8080", sc.logRequest(http.DefaultServeMux)))
	//sc.deviceManager.RunAngleTest()
}

func (sc *ServoCoolant) RegisterEndpoints() {
	http.HandleFunc("/", sc.handler)

}

func (sc *ServoCoolant) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I lo ve %s!", r.URL.Path[1:])
}

func (sc *ServoCoolant) logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sc.log.WithFields(log.Fields{
			"address": r.RemoteAddr,
			"method":  r.Method,
			"url":     r.URL,
		}).Debug("http request")
		handler.ServeHTTP(w, r)
	})
}
