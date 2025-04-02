package storage

import (
	"database/sql"
	"log"

	"app/models"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "postgres://postgres:postgres@db:5432/shoppingdb?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			quantity INTEGER NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("❌ Fehler beim Erstellen der Tabelle:", err)
	}
}

func GetAllItems() ([]models.Item, error) {
	rows, err := DB.Query("SELECT id, name, quantity FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
