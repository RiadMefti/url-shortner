package repository

import (
	"database/sql"
	"fmt"

	"github.com/RiadMefti/url-shortner/models"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	Db *sql.DB
}

func (r *Repository) PostUrl(uniqueUrl string, url string) error {
	query := "INSERT INTO urls(id_url, original_url) VALUES(?, ?)"
	stmt, err := r.Db.Prepare(query)
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

func (r *Repository) UrlExists(url string) (*models.URLExists, error) {
	query := "SELECT id_url FROM urls WHERE original_url = ?"
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		// Handle error
		return nil, err
	}
	defer stmt.Close()

	var originalUrl *string
	err = stmt.QueryRow(url).Scan(&originalUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
			fmt.Println("No rows found")
			return &models.URLExists{Exists: false, IdUrl: nil}, nil
		}

		return nil, err
	}

	fmt.Println("Original URL:", *originalUrl)
	return &models.URLExists{Exists: true, IdUrl: originalUrl}, nil
}

func (r *Repository) IDExists(url string) (bool, error) {
	query := "SELECT id_url FROM urls WHERE original_url = ?"
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		// Handle error
		return true, err
	}
	defer stmt.Close()

	var id *string
	err = stmt.QueryRow(url).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
			fmt.Println("No rows found")
			return false, nil
		}

		return true, err
	}

	return true, nil
}

func (r *Repository) GetUrl(id string) (string, error) {
	query := "SELECT original_url FROM urls WHERE id_url = ?"
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		// Handle error
		return "", err
	}
	defer stmt.Close()

	var url string
	err = stmt.QueryRow(id).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no rows were returned
			fmt.Println("No rows found")
			return "", nil
		}

		return "", err
	}

	return url, nil
}
