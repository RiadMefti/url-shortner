package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type MainPage struct {
	URL    string
	NewURL string
}

type ShortenedUrlPage struct {
	URL     string
	NewURL  string
	Success bool
	Message string
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, &MainPage{})
}

func main() {
	// Serve static files
	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route handlers
	http.HandleFunc("GET /", serveIndex)
	http.HandleFunc("POST /parse", parseForm)

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseForm(w http.ResponseWriter, r *http.Request) {

	inputURL := strings.TrimSpace(r.FormValue("url"))

	newUrl := inputURL + "HEHEHE NEW URL"

	data := ShortenedUrlPage{
		URL:     inputURL,
		NewURL:  newUrl,
		Success: true,
		Message: fmt.Sprintf("Successfully parsed URL: %s", inputURL),
	}

	tmpl, _ := template.ParseFiles("templates/url.html")
	tmpl.Execute(w, data)
}
