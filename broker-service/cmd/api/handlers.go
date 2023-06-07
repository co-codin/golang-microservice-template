package main

import "net/http"

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

}