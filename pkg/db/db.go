package db

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/snowflake-software/polarprint/pkg/types"
	"github.com/snowflake-software/polarprint/pkg/utils"
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
			quantity INT NOT NULL,
			clusterId INT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS clusters(
			id INT NOT NULL PRIMARY KEY,
			key TEXT NOT NULL,
			printers TEXT NOT NULL
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

func GetClusters() ([]types.Cluster, error) {
	rows, err := DB.Query("SELECT * FROM clusters")
	if err != nil {
		return nil, err
	}

	var clusters []types.Cluster

	for rows.Next() {
		var cluster types.Cluster
		var printersUnparsed string

		if err := rows.Scan(&cluster.Id, &cluster.Key, &printersUnparsed); err != nil {
			return nil, err
		}

		cluster.Printers = utils.UnpackPrinterArray(printersUnparsed)

		clusters = append(clusters, cluster)
	}

	return clusters, nil
}

func GetCluster(id int64) (*types.Cluster, error) {
	var cluster types.Cluster
	var printersUnparsed string

	if err := DB.QueryRow("SELECT * FROM clusters WHERE id = ?", id).Scan(&cluster.Id, &cluster.Key, &printersUnparsed); err != nil {
		return nil, err
	}

	cluster.Printers = utils.UnpackPrinterArray(printersUnparsed)

	return &cluster, nil
}

func CreateCluster() (*types.Cluster, error) {
	id := time.Now().Unix() + utils.RandomRange(11111111, 99999999)
	key := utils.RandomAlphanumberic(8)

	_, err := DB.Exec(`INSERT INTO clusters(id, key, printers) values(?, ?, ?)`,
		id,
		key,
		"",
	)

	if err != nil {
		return nil, err
	}

	cluster, err := GetCluster(id)

	if err != nil {
		log.Fatal("Cluster not found, even though just created", err)
	}

	return cluster, nil
}
