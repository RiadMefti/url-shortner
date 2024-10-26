package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func PostUrl(uniqueUrl string, url string) error {
	db, err := sql.Open("sqlite3", "app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "INSERT INTO urls(id_url, original_url) VALUES(?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, dberr := stmt.Exec(uniqueUrl, url)
	if dberr != nil {
		return dberr
	}

	return nil
}
