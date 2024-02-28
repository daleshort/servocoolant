package main

import (
	"encoding/json"
	"net/http"
)

type ToolQueueAddRequest struct {
	ToolId int `json:"toolid" example:"2"`
}

func (sc *ServoCoolant) handlerAutoStart(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		sc.autoManager.HandleProgramStartEvent()
	}
}

func (sc *ServoCoolant) handlerAutoEnd(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		sc.autoManager.HandleEndOfProgramEvent()
	}
}

func (sc *ServoCoolant) handlerToolQueueAdd(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var req ToolQueueAddRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			sc.log.Error("bad post tool to queue request")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sc.autoManager.AddToolToQueue(req.ToolId)
	}
}
