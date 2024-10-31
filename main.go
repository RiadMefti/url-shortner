package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/RiadMefti/url-shortner/filehandler"
	"github.com/RiadMefti/url-shortner/repository"
	"github.com/RiadMefti/url-shortner/services"
	_ "github.com/mattn/go-sqlite3"
)

func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Internal server error: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func initializeDB() (*sql.DB, error) {
	// Create a data directory in /tmp for Railway

	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Initialize your database schema here if needed
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id_url TEXT PRIMARY KEY,
			original_url TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Initialize the database connection with error handling
	db, err := initializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
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

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Health check endpoint for Railway
	mux.Handle("GET /health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	// Serve static files
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route handlers with error handling middleware
	mux.Handle("GET /", errorHandler(http.HandlerFunc(fileHandler.ServeIndex)))
	mux.Handle("POST /parse", errorHandler(http.HandlerFunc(fileHandler.ParseForm)))
	mux.Handle("GET /{id}", errorHandler(http.HandlerFunc(urlService.HandleRedirect)))

	// Use `PORT` provided in environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
