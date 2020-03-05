package controllers

import (
	"Fit-Time-Backend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()
	fmt.Println(params)
	// Create a new model
	hello := models.Hello{
		FirstName: params.Get("Name"),
		LastName:  "Arnall",
	}

	// Encode json from hello onto the interface w, the responsewriter
	jsonString, err := json.Marshal(hello)
	if err != nil {
		log.Println("Error marshalling json: ", err.Error())
	}

	w.Header().Add("Content-Type", "application/json")

	_, _ = w.Write(jsonString)
}
