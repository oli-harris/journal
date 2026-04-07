package httpapi

import (
	"net/http"

	handlers "journal/internal/api/handlers"

	"github.com/uptrace/bun"
)

func NewRouter(db *bun.DB) http.Handler {
	h := handlers.New(db)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", h.HelloWorld)

	return mux
}
