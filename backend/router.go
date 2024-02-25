package main

import (
	"net/http"
	"fmt"
)

func (sc *ServoCoolant) RegisterEndpoints() {
	http.HandleFunc("/", sc.handler)
	
}

func (sc *ServoCoolant) handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}