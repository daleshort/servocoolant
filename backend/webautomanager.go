package main

import (
	"encoding/json"
	"net/http"
)

type ToolQueueRequest struct {
	ToolId int `json:"toolid" example:"2"`
}

func (sc *ServoCoolant) handlerAutoStart(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		sc.autoManager.HandleProgramStartEvent()
		return
	}
	http.Error(w, "invalid method", http.StatusMethodNotAllowed)

}

func (sc *ServoCoolant) handlerAutoEnd(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		sc.autoManager.HandleEndOfProgramEvent()
		return
	}
	http.Error(w, "invalid method", http.StatusMethodNotAllowed)

}

func (sc *ServoCoolant) handlerToolQueueAdd(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var req ToolQueueRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			sc.log.Error("bad post tool to queue request")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sc.autoManager.AddToolToQueue(req.ToolId)
		return
	}
	http.Error(w, "invalid method", http.StatusMethodNotAllowed)

}

func (sc *ServoCoolant) handlerToolQueuePosition(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var req ToolQueueRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			sc.log.Error("bad post tool to queue request")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = sc.autoManager.HandleSetToolQueueToPosition(req.ToolId)

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			sc.log.Error(err.Error())
		}
		return
	}

	
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		sc.log.Error("invalid method requested")
}
