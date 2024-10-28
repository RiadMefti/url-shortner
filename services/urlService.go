package services

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/RiadMefti/url-shortner/models"
	"github.com/RiadMefti/url-shortner/repository"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func CreateURl(url string) string {

	exists := urlExists(url)
	if exists.Exists {
		return *exists.IdUrl
	}
	uniqueUrl := createUniqueId()
	err := repository.PostUrl(uniqueUrl, url)

	if err != nil {
		log.Fatal(err)
	}

	return uniqueUrl

}

func urlExists(url string) *models.URLExists {

	exists, err := repository.UrlExists(url)

	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func idExists(id string) bool {

	exists, err := repository.IDExists(id)

	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func createUniqueId() string {
	idLength := 8
	uniqueId := make([]byte, idLength)
	for i := range uniqueId {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		uniqueId[i] = charset[randomIndex.Int64()]
	}
	if idExists(string(uniqueId)) {
		createUniqueId()
	}
	return string(uniqueId)
}
