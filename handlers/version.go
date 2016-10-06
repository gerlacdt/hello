package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

var version string = "v0.1.0"

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{
		Version: version,
	}
	log.Printf("/version request --> %s\n", version)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
