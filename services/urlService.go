package services

import (
	"crypto/rand"
	"log"
	"math/big"
	"net/http"

	"github.com/RiadMefti/url-shortner/models"
	"github.com/RiadMefti/url-shortner/repository"
)

type UrlService struct {
	Repository *repository.Repository
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (s *UrlService) CreateURl(url string) string {

	exists := s.urlExists(url)
	if exists.Exists {
		return *exists.IdUrl
	}
	uniqueUrl := s.createUniqueId()
	err := s.Repository.PostUrl(uniqueUrl, url)

	if err != nil {
		log.Fatal(err)
	}

	return uniqueUrl

}

func (s *UrlService) urlExists(url string) *models.URLExists {

	exists, err := s.Repository.UrlExists(url)

	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func (s *UrlService) idExists(id string) bool {

	exists, err := s.Repository.IDExists(id)

	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func (s *UrlService) createUniqueId() string {
	idLength := 8
	uniqueId := make([]byte, idLength)
	for i := range uniqueId {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		uniqueId[i] = charset[randomIndex.Int64()]
	}
	if s.idExists(string(uniqueId)) {
		s.createUniqueId()
	}
	return string(uniqueId)
}

func (s *UrlService) GetUrl(id string) (string, error) {
	url, err := s.Repository.GetUrl(id)
	if err != nil {
		return "", err
	}
	return url, nil
}
func (s *UrlService) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL path
	id := r.PathValue("id")
	url, err := s.GetUrl(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)

}
