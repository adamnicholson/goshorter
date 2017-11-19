package shortner

import (
	"database/sql"
	"log"
)

type Repository struct {
	db *sql.DB
}

func (r Repository) addShortUrl (alias string, long string) {
	_, err := r.db.Query(`INSERT INTO short_url VALUES (?, ?, NOW());`, alias, long)
	if err != nil {
		log.Fatalf("Error inserting row: %s", err.Error())
	}
}

func (r Repository) getLocation (alias string) (string, error) {

	var location string
	rows := r.db.QueryRow(`SELECT location FROM short_url WHERE alias = ?;`, alias)

	err := rows.Scan(&location)

	return location, err
}