/*
	Simple Rest Service example without Database. I recommend to use 'dep' to install dependencies
	ch.rajshekar@gmail.com
 */

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"restservice/controllers"
)

func main() {

	router := mux.NewRouter()

	// GET METHOD
	router.HandleFunc("/api/sayhello",controllers.SayHi).Methods(http.MethodGet)

	// POST METHOD
	router.HandleFunc("/api/sayhello", controllers.PostHi).Methods(http.MethodPost)


	port := "8000"
	fmt.Println("Running on port: " + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
