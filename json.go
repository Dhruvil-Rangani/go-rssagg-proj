package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Server error: %s", msg)
	} 
	type errorResponse struct {
		Error string `json:"error"`
	}

	jsonResponse(w, status, errorResponse{Error: msg});
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	dat, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling JSON data of type %T: %v", data, err)
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(dat)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}
}