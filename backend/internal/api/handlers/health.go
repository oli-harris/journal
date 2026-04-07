package handlers

import (
	"encoding/json"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func HelloWorld(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, helloWorldResponse{Message: "Hello, World!"})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
