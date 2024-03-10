package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServoStatusResponse struct {
	ServoStatus map[int]ServoDetailStatusResponse `json:"servostatus" example:"100"`
}

type ServoDetailStatusResponse struct {
	Angle        int     `json:"angle" example:"100"`
	IsAuto       bool    `json:"isauto" example:"true"`
	IsWiggle     bool    `json:"iswiggle" example:"true"`
	Amplitude    int     `json:"amplitude" example:"10"`
	Frequency    float32 `json:"frequency" example:".5"`
	TravelRange  int     `json:"travelrange" example:"265"`
	Offset       int     `json:"offset" example:"20"`
	SoftLimitMin int     `json:"softlimitmin" example:"0"`
	SoftLimitMax int     `json:"softlimitmax" example:"100"`
}

type ServoPostRequest struct {
	Servos []int `json:"servos" example:"[1,2]"`
	Angle  int   `json:"angle" example:"100"`
}

type ServoWigglePostRequest struct {
	Servos    []int    `json:"servos" example:"[1,2]"`
	Amplitude *int     `json:"amplitude" example:"10"`
	Frequency *float32 `json:"frequency" example:".5"`
	IsWiggle  *bool    `json:"iswiggle" example:"true"`
}
type ServoAutoPostRequest struct {
	Servos []int `json:"servos" example:"[1,2]"`
	IsAuto bool  `json:"isauto" example:"true"`
}

func (sc *ServoCoolant) handlerServoWiggle(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		sc.handlerPostServoWiggle(w, r)
		return
	}

}

func (sc *ServoCoolant) handlerPostServoWiggle(w http.ResponseWriter, r *http.Request) {

	var req ServoWigglePostRequest

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
			if req.IsWiggle != nil {
				sc.deviceManager.Servo1.IsWiggle = *req.IsWiggle
			}
			if req.Amplitude != nil {
				sc.deviceManager.Servo1.WiggleAmplitude = *req.Amplitude
			}
			if req.Frequency != nil {
				sc.deviceManager.Servo1.WiggleFrequency = *req.Frequency
			}
		} else if servo == 2 {
			if req.IsWiggle != nil {
				sc.deviceManager.Servo2.IsWiggle = *req.IsWiggle
			}
			if req.Amplitude != nil {
				sc.deviceManager.Servo2.WiggleAmplitude = *req.Amplitude
			}
			if req.Frequency != nil {
				sc.deviceManager.Servo2.WiggleFrequency = *req.Frequency
			}
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

func (sc *ServoCoolant) handlerServoAuto(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		sc.handlerPostServoAuto(w, r)
		return
	}

}

func (sc *ServoCoolant) handlerPostServoAuto(w http.ResponseWriter, r *http.Request) {

	var req ServoAutoPostRequest

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
			sc.deviceManager.Servo1.IsAuto = req.IsAuto
		} else if servo == 2 {
			sc.deviceManager.Servo2.IsAuto = req.IsAuto
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

func (sc *ServoCoolant) getServoStatus() map[int]ServoDetailStatusResponse {
	return map[int]ServoDetailStatusResponse{
		1: ServoDetailStatusResponse{
			Angle:       sc.deviceManager.Servo1.Angle,
			IsAuto:      sc.deviceManager.Servo1.IsAuto,
			IsWiggle:    sc.deviceManager.Servo1.IsWiggle,
			Amplitude:   sc.deviceManager.Servo1.WiggleAmplitude,
			Frequency:   sc.deviceManager.Servo1.WiggleFrequency,
			TravelRange: sc.deviceManager.Servo1.TravelRange,
			Offset:      sc.deviceManager.Servo1.Offset,
			SoftLimitMin: sc.deviceManager.Servo1.SoftLimitMin,
			SoftLimitMax: sc.deviceManager.Servo1.SoftLimitMax,
		},
		2: ServoDetailStatusResponse{
			Angle:       sc.deviceManager.Servo2.Angle,
			IsAuto:      sc.deviceManager.Servo2.IsAuto,
			IsWiggle:    sc.deviceManager.Servo2.IsWiggle,
			Amplitude:   sc.deviceManager.Servo2.WiggleAmplitude,
			Frequency:   sc.deviceManager.Servo2.WiggleFrequency,
			TravelRange: sc.deviceManager.Servo2.TravelRange,
			Offset:      sc.deviceManager.Servo2.Offset,
			SoftLimitMin: sc.deviceManager.Servo2.SoftLimitMin,
			SoftLimitMax: sc.deviceManager.Servo2.SoftLimitMax,
		},
	}
}

func (sc *ServoCoolant) handlerServo(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		sc.handlerGetServo(w, r)
		return
	} else if r.Method == http.MethodPost {
		sc.handlerPostServo(w, r)
		return
	}

}

func (sc *ServoCoolant) handlerGetServo(w http.ResponseWriter, r *http.Request) {

	resp := ServoStatusResponse{ServoStatus: sc.getServoStatus()}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (sc *ServoCoolant) handlerPostServo(w http.ResponseWriter, r *http.Request) {

	var req ServoPostRequest

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
