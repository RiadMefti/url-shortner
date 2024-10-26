package services

import (
	"log"

	"github.com/RiadMefti/url-shortner/repository"
)

func CreateURl() {
	err := repository.PostUrl("123345", "google.conm")

	if err != nil {
		log.Fatal(err)
	}

	print("GOOD")
}
