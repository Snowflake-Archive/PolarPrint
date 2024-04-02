package db

import (
	"database/sql"
	"log"
)

func SetupDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./polarprint.db")

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS queue(
			id INT NOT NULL PRIMARY KEY,
			file TEXT NOT NULL,
			quantity INT NOT NULL
		);
	`)

	if err != nil {
		return nil, err
	}

	log.Print("Successfully initiated database!")

	count := 0
	err = db.QueryRow("SELECT COUNT(*) FROM queue").Scan(&count)

	if err != nil {
		return nil, err
	}

	log.Printf("There are %d queued items.", count)

	return db, nil
}
