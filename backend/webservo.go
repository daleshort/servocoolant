package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type servoStatusResponse struct {
	Servo1angle int `json:"servo1angle" example:"100"`
	Servo2angle int `json:"servo2angle" example:"100"`
}

type servoPostRequest struct {
	Servos []int `json:"servos" example:"[1,2]"`
	Angle  int   `json:"angle" example:"100"`
}

func (sc *ServoCoolant) handlerServo(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		sc.handlerGetServo(w, r)
	} else if r.Method == http.MethodPost {
		sc.handlerPostServo(w, r)
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

func (sc *ServoCoolant) handlerPostServo(w http.ResponseWriter, r *http.Request) {

	var req servoPostRequest

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sc.log.Error("bad post servo request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, servo := range req.Servos {
		var err error
		if servo == 1 {
			err = sc.deviceManager.Servo1.SetAngle(req.Angle)
		} else if servo == 2 {
			err = sc.deviceManager.Servo2.SetAngle(req.Angle)
		} else {
			err = fmt.Errorf("bad servo number requested %v", servo)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			sc.log.Error(err.Error())
		}
	}
	// returns response HTTP 200 OK by default
}
