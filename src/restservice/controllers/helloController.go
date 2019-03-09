package controllers

import (
		"net/http"
		"restservice/models"
		u "restservice/utils"
	"encoding/json"
)

// Say hello call handler
// Will generate JSON object which contains a message and actual response

var SayHi = func(w http.ResponseWriter, r *http.Request){
	helloObj := &models.Account{"Rajshekar","+91-9922616111","ch.rajshekar@gmail.com"}
	resp := helloObj.SayHi()
	u.Respond(w, resp)
}

// Post hi call handler
// Will post back JSON object with a message and JSON sent to API

var PostHi = func(w http.ResponseWriter, r *http.Request){
	helloObj := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(helloObj) // decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := helloObj.SayHi() // Say hi with posted object
	u.Respond(w, resp)
}