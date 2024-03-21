package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", echoHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
    // Read request body
    decoder := json.NewDecoder(r.Body)
    var requestBody interface{}
    err := decoder.Decode(&requestBody)
    if err != nil {
        http.Error(w, "Error decoding request body", http.StatusBadRequest)
        return
    }

    // Set Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Encode the request body as JSON and write it to the response
    encoder := json.NewEncoder(w)
    err = encoder.Encode(requestBody)
    if err != nil {
        log.Printf("Error encoding response: %v", err)
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
    }
}
