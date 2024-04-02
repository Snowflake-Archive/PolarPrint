package db

import (
	"database/sql"
	"log"
)

var (
	DB *sql.DB
)

type QueueItem struct {
	Id       int    `json:"id"`
	File     string `json:"file"`
	Quantity int    `json:"quantity"`
}

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

	DB = db
	return db, nil
}

func GetQueue() ([]QueueItem, error) {
	rows, err := DB.Query("SELECT * FROM queue")
	if err != nil {
		return nil, err
	}

	var orders []QueueItem

	for rows.Next() {
		var order QueueItem

		if err := rows.Scan(&order.Id, &order.File, &order.Quantity); err != nil {
			return nil, err
		}

		log.Print(order)
		orders = append(orders, order)
	}

	return orders, nil
}
