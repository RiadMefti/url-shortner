package filehandler

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/RiadMefti/url-shortner/services"
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

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, &MainPage{})
}

func ParseForm(w http.ResponseWriter, r *http.Request) {

	inputURL := strings.TrimSpace(r.FormValue("url"))

	services.CreateURl()

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
