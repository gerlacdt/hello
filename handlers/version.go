package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

var Version string = "v0.5.0"

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{
		Version: Version,
	}
	log.Printf("/version request --> %s\n", Version)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
