package filehandler

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/RiadMefti/url-shortner/models"

	"github.com/RiadMefti/url-shortner/services"
)

type StaticFile struct {
	UrlService *services.UrlService
}

func (s *StaticFile) ServeIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, &models.MainPage{})
}

func (s *StaticFile) ParseForm(w http.ResponseWriter, r *http.Request) {
	inputURL := strings.TrimSpace(r.FormValue("url"))
	id := s.UrlService.CreateURl(inputURL)

	// Get the scheme (http/https)
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	// If you're behind a proxy that sets X-Forwarded-Proto header
	if forwardedProto := r.Header.Get("X-Forwarded-Proto"); forwardedProto != "" {
		scheme = forwardedProto
	}

	// Get the host
	host := r.Host

	// Construct the base URL
	baseURL := fmt.Sprintf("%s://%s", scheme, host)

	// Create the new shortened URL
	newUrl := fmt.Sprintf("%s/%s", baseURL, id)

	data := models.ShortenedUrlPage{
		URL:     inputURL,
		NewURL:  newUrl,
		Success: true,
		Message: fmt.Sprintf("Successfully parsed URL: %s", inputURL),
	}

	tmpl, _ := template.ParseFiles("templates/url.html")
	tmpl.Execute(w, data)
}
