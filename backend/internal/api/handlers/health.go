package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/uptrace/bun"
)

type Handler struct {
	db *bun.DB
}

func New(db *bun.DB) *Handler {
	return &Handler{db: db}
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func (h *Handler) HelloWorld(w http.ResponseWriter, _ *http.Request) {
	_ = h.db
	writeJSON(w, http.StatusOK, helloWorldResponse{Message: "Hello, World!"})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
