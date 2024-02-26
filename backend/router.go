package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (sc *ServoCoolant) RegisterEndpoints() {
	http.HandleFunc("/servo", sc.handlerServo)
	http.HandleFunc("/", sc.handlerTest)

}

func (sc *ServoCoolant) handlerTest(w http.ResponseWriter, r *http.Request) {
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

func (sc *ServoCoolant) enableCors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		
		handler.ServeHTTP(w, r)
	})
}



// fs := http.FileServer(http.Dir("build"))
// http.Handle("/", fs)
