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

}

func (sc *ServoCoolant) handlerAutoEnd(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		sc.autoManager.HandleEndOfProgramEvent()
		return
	}

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

}
func (sc *ServoCoolant) handlerForceTool(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var req ToolQueueRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			sc.log.Error("bad post tool to force tool request")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sc.autoManager.HandleEndOfProgramEvent()
		sc.autoManager.AddToolToQueue(req.ToolId)
		sc.autoManager.HandleProgramStartEvent()
		sc.autoManager.HandleSetToolQueueToPosition(0)


		err = sc.autoManager.HandleSetToolQueueToPosition(req.ToolId)

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			sc.log.Error(err.Error())
		}
		return
	}

}
