package main

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (sc *ServoCoolant) RegisterEndpoints() {
	http.HandleFunc("/auto/toolqueueadd", sc.handlerToolQueueAdd)
	http.HandleFunc("/auto/programstart", sc.handlerAutoStart)
	http.HandleFunc("/auto/programend", sc.handlerAutoEnd)
	http.HandleFunc("/servo", sc.handlerServo)
	http.HandleFunc("/servowiggle", sc.handlerServoWiggle)
	http.HandleFunc("/servoauto", sc.handlerServoAuto)
	http.HandleFunc("/status", sc.handlerStatus)
	http.HandleFunc("/", sc.handlerTest)

}

func (sc *ServoCoolant) handlerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I lo ve %s!", r.URL.Path[1:])
}

func (sc *ServoCoolant) logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if(!strings.Contains(r.URL.Path,"/status")){
		sc.log.WithFields(log.Fields{
			"address": r.RemoteAddr,
			"method":  r.Method,
			"url":     r.URL,
		}).Debug("http request")
		}
		
		handler.ServeHTTP(w, r)
	})
}

func (sc *ServoCoolant) enableCors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	
		handler.ServeHTTP(w, r)
	})
}



// fs := http.FileServer(http.Dir("build"))
// http.Handle("/", fs)
