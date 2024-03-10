package main

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var BUILD_PATH =  "../servofrontend/dist"
var ASSET_PATH =  "../servofrontend/dist/assets"

func (sc *ServoCoolant) RegisterEndpoints() {

	//not sure if this is needed?
	assetHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir(ASSET_PATH)))
	
	http.HandleFunc("/auto/toolqueueadd", sc.handlerToolQueueAdd)
	http.HandleFunc("/auto/programstart", sc.handlerAutoStart)
	http.HandleFunc("/auto/queueposition", sc.handlerToolQueuePosition)
	http.HandleFunc("/auto/forcetool", sc.handlerForceTool)
	http.HandleFunc("/auto/programend", sc.handlerAutoEnd)
	http.HandleFunc("/servo", sc.handlerServo)
	http.HandleFunc("/servowiggle", sc.handlerServoWiggle)
	http.HandleFunc("/servoauto", sc.handlerServoAuto)
	http.HandleFunc("/status", sc.handlerStatus)
	http.HandleFunc("/toollength", sc.handlerPostToolLength)
	http.Handle("/assets/", assetHandler)
	http.Handle("/",http.FileServer(http.Dir(BUILD_PATH)))

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
