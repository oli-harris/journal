package web

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed all:static/*
var files embed.FS

func SvelteKitHandler(path string) http.Handler {
	fsys, err := fs.Sub(files, "static")
	if err != nil {
		log.Fatal(err)
	}
	filesystem := http.FS(fsys)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, path)
		// try if file exists at path, if not append .html (SvelteKit adapter-static specific)
		_, err := filesystem.Open(path)
		if errors.Is(err, os.ErrNotExist) {
			path = fmt.Sprintf("%s.html", path)
		}
		r.URL.Path = path
		http.FileServer(filesystem).ServeHTTP(w, r)
	})
}
