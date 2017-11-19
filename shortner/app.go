package shortner

import (
	"database/sql"
	"log"
)

type Container struct {
	db         *sql.DB
	repository Repository
}

func (app *Container) Boot() *Container {

	var err error
	app.db, err = sql.Open("mysql", "root:@/goshorter")
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	app.db.Begin()
	app.repository = Repository{db: app.db}

	resetDatabase(app.db)

	return app
}

func resetDatabase(db *sql.DB) {
	db.Query("DROP TABLE short_url;")
	_, err := db.Query("CREATE TABLE short_url (alias VARCHAR(255), location VARCHAR(255), created_at DATETIME);")
	if err != nil {
		log.Fatalf("Failed to run database migration(s): %s", err.Error())
	}
}