package main

import (
	"encoding/json"
	"net/http"
)

type servoStatusResponse struct {
	Servo1angle int `json:"servo1angle" example:"100"`
	Servo2angle int `json:"servo2angle" example:"100"`
}

func (sc *ServoCoolant) handlerServo(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		sc.handlerGetServo(w, r)
	}

}

func (sc *ServoCoolant) handlerGetServo(w http.ResponseWriter, r *http.Request) {

	resp := servoStatusResponse{
		Servo1angle: sc.deviceManager.Servo1.Angle,
		Servo2angle: sc.deviceManager.Servo2.Angle,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
