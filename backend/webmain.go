package main

import (
	"encoding/json"
	"net/http"

	config "mechied.com/servocoolant/config"
)

type StatusResponse struct {
	ServoStatus     map[int]ServoDetailStatusResponse `json:"servostatus" example:"100"`
	IsToolsenseHigh bool                              `json:"istoolsensehigh" example:"true" `
	Tools           map[int]config.Tool               `json:"tools" example:"{1:{length:12.2}}" `
	ToolQueue       []int                             `json:"toolqueue" example:"[2,1,12]" `
	IsProgramRunning         bool `json:"isprogramrunning" example:"true" `
	CurrentToolQueuePosition int  `json:"currenttoolqueueposition" example:"1" `
}

type ToolLengthRequest struct {
	ToolId int `json:"toolid" example:"2"`
	ToolLength float32 `json:"toollength" example:"1.43"`
}

func (sc *ServoCoolant) handlerStatus(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tools, _ := sc.config.GetAllToolLengths()
		resp := StatusResponse{
			ServoStatus:              sc.getServoStatus(),
			IsToolsenseHigh:          sc.deviceManager.IsToolsenseHigh,
			Tools:                    tools,
			ToolQueue:                sc.autoManager.ToolQueue,
			IsProgramRunning:         sc.autoManager.IsProgramRunning,
			CurrentToolQueuePosition: sc.autoManager.CurrentToolQueuePosition,

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return

	}

}

func (sc *ServoCoolant) handlerPostToolLength(w http.ResponseWriter, r *http.Request) {

		if(r.Method == http.MethodPost){

			var req ToolLengthRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				sc.log.Error("bad post tool length request")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		
			sc.config.SetToolLength(req.ToolId,req.ToolLength)

		}

}