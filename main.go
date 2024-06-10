package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRequest)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	query := r.URL.Query()

	fmt.Println("")
	log.Println("Received request for:", path, "with query:", query)

	if !strings.HasPrefix(path, "/github.com/aarshkshah1992/prebuilt-ffi-darwin-arm64") {
		// redirect request to google go proxy
		http.Redirect(w, r, "https://proxy.golang.org"+path, http.StatusMovedPermanently)
	}

	// Standard module handling below this point
	parts := strings.Split(path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid module request", http.StatusBadRequest)
		return
	}
	version := parts[len(parts)-1]

	switch {
	case strings.HasSuffix(version, "list"):
		http.Redirect(w, r, "https://proxy.golang.org"+path, http.StatusMovedPermanently)
	case strings.HasSuffix(version, ".info"):
		http.Redirect(w, r, "https://proxy.golang.org"+path, http.StatusMovedPermanently)
	case strings.HasSuffix(version, ".mod"):
		http.Redirect(w, r, "https://proxy.golang.org"+path, http.StatusMovedPermanently)
	case strings.HasSuffix(version, ".zip"):
		handleZip(w, version)
	default:
		http.Error(w, "Unsupported module request", http.StatusNotImplemented)
	}
}
func handleZip(w http.ResponseWriter, version string) {
	zipFilePath := "/Users/aarshshah/src/filecoin/prebuilt-ffi-poc/module.zip" // Update this path accordingly
	zipData, err := ioutil.ReadFile(zipFilePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Write(zipData)
	fmt.Println("Served zip file for version:", version)
}
