package main

import (
	"log"
	"net/http"

	"github.com/RiadMefti/url-shortner/filehandler"
)

func main() {

	// Serve static files
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Route handlers
	http.HandleFunc("GET /", filehandler.ServeIndex)
	http.HandleFunc("POST /parse", filehandler.ParseForm)

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
