package main

import (
	"net/http"
)

type statusResponse struct {
	ServoStatus map[int]ServoDetailStatusResponse `json:"servostatus" example:"100"`
}

func (sc *ServoCoolant) handlerStatus(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		
	}

}
