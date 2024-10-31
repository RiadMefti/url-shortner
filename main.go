package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/RiadMefti/url-shortner/filehandler"
	"github.com/RiadMefti/url-shortner/repository"
	"github.com/RiadMefti/url-shortner/services"
)

func main() {
	// Initialize the database connection
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the repository with the database connection
	repo := repository.Repository{
		Db: db,
	}

	urlService := services.UrlService{
		Repository: &repo,
	}
	fileHandler := filehandler.StaticFile{
		UrlService: &urlService,
	}

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route handlers
	http.HandleFunc("/", fileHandler.ServeIndex)
	http.HandleFunc("/parse", fileHandler.ParseForm)

	// Add the dynamic route handler
	http.HandleFunc("/{id}", urlService.HandleRedirect)

	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
