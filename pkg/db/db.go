package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/snowflake-software/polarprint/pkg/types"
)

var (
	DB *sql.DB
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

	DB = db
	return db, nil
}

func GetQueue() ([]types.QueueItem, error) {
	rows, err := DB.Query("SELECT * FROM queue")
	if err != nil {
		return nil, err
	}

	var orders []types.QueueItem

	for rows.Next() {
		var order types.QueueItem

		if err := rows.Scan(&order.Id, &order.File, &order.Quantity); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func InsertOrder(id int64, data types.QueuePrintRequestBody) error {
	_, err := DB.Exec(`INSERT INTO queue(id, file, quantity) values(?, ?, ?)`,
		id,
		"./prints/"+data.FilePath,
		data.Quantity,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteOrder(id int64) error {
	res, err := DB.Exec("DELETE FROM queue WHERE id = ?", id)

	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("order not found")
	}

	return nil
}
