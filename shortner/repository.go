
package shortner

import (
	"database/sql"
	"log"
)

type Repository struct {
	db *sql.DB
}

type Shorten struct {
	short string
	long string
}

func (r Repository) addShortUrl (alias string, long string) {
	_, err := r.db.Query(`INSERT INTO short_url VALUES (?, ?, NOW());`, alias, long)
	if err != nil {
		log.Fatalf("Error inserting row: %s", err.Error())
	}
}

func (r Repository) get() []Shorten {

	var all []Shorten

	rows, _ := r.db.Query(`SELECT alias, location FROM short_url;`)

	for rows.Next() {
		var shorten Shorten

		err := rows.Scan(&shorten.short, &shorten.long)

		if err != nil {
			log.Fatalln(err)
		}

		all = append(all, shorten)
	}

	return all
}

func (r Repository) getLocation (alias string) (string, error) {

	var location string
	rows := r.db.QueryRow(`SELECT location FROM short_url WHERE alias = ?;`, alias)

	err := rows.Scan(&location)

	return location, err
}