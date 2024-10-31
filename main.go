package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/RiadMefti/url-shortner/repository"
	"github.com/RiadMefti/url-shortner/services"
	_ "github.com/mattn/go-sqlite3"
)

// errorHandler is a middleware that catches errors and returns a 502 status code
func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Internal server error: %v", err)
				http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

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
	fileHandler := services.StaticFileService{
		UrlService: &urlService,
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route handlers with error handling middleware
	mux.Handle("GET /", errorHandler(http.HandlerFunc(fileHandler.ServeIndex)))
	mux.Handle("POST /parse", errorHandler(http.HandlerFunc(fileHandler.ParseForm)))
	mux.Handle("GET /{id}", errorHandler(http.HandlerFunc(urlService.HandleRedirect)))

	// Use `PORT` provided in environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
