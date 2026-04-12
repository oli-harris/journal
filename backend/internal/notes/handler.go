package notes

import (
	"encoding/json"
	"net/http"
	"strings"
)

type NotesHandler struct {
	repo Repository
}

func NewNotesHandler(repo Repository) *NotesHandler {
	return &NotesHandler{repo: repo}
}

func (h *NotesHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.repo.List(r.Context(), 1)

	if err != nil {
		http.Error(w, "error fetching notes", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, notes)
}

type createNoteRequest struct {
	Note string `json:"note"`
}

func (h *NotesHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req createNoteRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	req.Note = strings.TrimSpace(req.Note)
	if req.Note == "" {
		http.Error(w, "note is required", http.StatusBadRequest)
		return
	}

	uuid, err := h.repo.Create(r.Context(), 1, req.Note)
	if err != nil {
		http.Error(w, "error creating note", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{"uuid": uuid})
}

func writeJSON(w http.ResponseWriter, statusCode int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(value)
}
