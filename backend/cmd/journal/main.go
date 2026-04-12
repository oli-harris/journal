package main

import (
	"journal/internal"
	"journal/web"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize database
	db, err := internal.InitialiseDB("db.sqlite")
	if err != nil {
		log.Fatalf("database init failed: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	// Handle API routes
	apiRouter := internal.NewRouter(db)
	mux.Handle("/api/", http.StripPrefix("/api", apiRouter))

	// Embedded frontend
	embeddedFrontendHandler := web.SvelteKitHandler("")
	mux.Handle("/", embeddedFrontendHandler)
	handler := mux

	log.Printf("Server running at http://localhost:%s", port)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
