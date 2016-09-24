package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "hello world",
	}
	log.Printf("/hello requested...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
