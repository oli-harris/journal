package internal

import (
	"database/sql"
	"journal/internal/notes"
	"net/http"
)

func NewRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	notesRepository := notes.NewSQLRepository(db)
	notesHandler := notes.NewNotesHandler(notesRepository)

	// Note routes
	mux.HandleFunc("GET /notes", notesHandler.GetNotes)
	mux.HandleFunc("POST /notes", notesHandler.CreateNote)

	return mux
}
