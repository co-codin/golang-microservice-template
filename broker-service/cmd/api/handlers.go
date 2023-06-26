package main

import (
	"net/http"
)

type RequestPayload struct {
	Action string `json:"action"`
	Auth AuthPayoload `json:"auth,omitempty"`
}

type AuthPayoload struct {

}


func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {

}