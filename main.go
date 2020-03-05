package main

import (
	"Fit-Time-Backend/controllers"
	"log"
	"net/http"
)

func main() {

	// List of Routes
	http.HandleFunc("/", controllers.Hello)

	// Starts the http server
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
