package models

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
