package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct for the response
	type Response struct {
		Message string `json:"message"`
	}

	// Create a response object
	response := Response{
		Message: "Hello World",
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the response object as JSON and write it to the response
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
