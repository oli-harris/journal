package main

import (
	"log"
	"net/http"
	"os"

	api "journal/internal/api"
	web "journal/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	// Handle API routes
	apiRouter := api.NewRouter()
	mux.Handle("/api/", http.StripPrefix("/api", apiRouter))

	// Embedded frontend
	embeddedFrontendHandler := web.SvelteKitHandler("")
	mux.Handle("/", embeddedFrontendHandler)
	handler := mux

	log.Printf("Server running at http://localhost:%s", port)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
