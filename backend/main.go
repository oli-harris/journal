package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func main() {
	devMode := false
	flag.BoolVar(&devMode, "dev", devMode, "enable dev mode")
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	// API routes are always enabled.
	mux.HandleFunc("/api/version", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"version":"1.0.0"}`))
	})

	mux.HandleFunc("/api/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"status":"ok","time":"%s"}`,
			time.Now().UTC().Format(time.RFC3339))))
	})

	var handler http.Handler = mux

	if devMode {
		handler = cors(handler)
		log.Printf("server running in dev mode on http://localhost:%s", port)
		log.Printf("embedded frontend disabled in dev mode; serving API routes only")
	} else {
		embeddedFrontendHandler := SvelteKitHandler("")
		// Only serve frontend from / in non-dev mode.
		mux.Handle("/", embeddedFrontendHandler)
		log.Printf("server running in prod mode on http://localhost:%s", port)
		log.Printf("embedded frontend handler mounted at /")
	}

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
