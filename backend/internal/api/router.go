package httpapi

import (
	"net/http"

	handlers "journal/internal/api/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HelloWorld)

	return mux
}
