package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func postUrl(uniqueUrl string, url string) error {
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}
	formattedQuery := fmt.Sprintf("INSERT INTO urls(id_url,original_url) VALUES(%s,%s)", uniqueUrl, url)

	_, dberr := db.Query(formattedQuery)

	if dberr != nil {
		return err
	}

	return nil

}
