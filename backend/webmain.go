package main

import (
	"encoding/json"
	"net/http"

	
)

type StatusResponse struct {
	ServoStatus     map[int]ServoDetailStatusResponse `json:"servostatus" example:"100"`
	IsToolsenseHigh bool                              `json:"istoolsensehigh" example:"true" `
	Tools           map[string]string               `json:"tools" example:"{1: "12.2"}" `
}

func (sc *ServoCoolant) handlerStatus(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tools, _ := sc.config.GetAllToolLengths()
		resp := StatusResponse{
			ServoStatus:     sc.getServoStatus(),
			IsToolsenseHigh: sc.deviceManager.IsToolsenseHigh,
			Tools:           tools,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	}

}
