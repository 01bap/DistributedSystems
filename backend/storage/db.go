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


func CreateOrUpdateItem(name string, quantity int) (models.Item, bool, error) {
	var item models.Item

	// Prüfe, ob ein Item mit diesem Namen schon existiert
	err := DB.QueryRow("SELECT id, quantity FROM items WHERE name = $1", name).Scan(&item.ID, &item.Quantity)

	if err == sql.ErrNoRows {
		// Neues Item anlegen
		err = DB.QueryRow("INSERT INTO items (name, quantity) VALUES ($1, $2) RETURNING id", name, quantity).
			Scan(&item.ID)
		if err != nil {
			return item, false, err
		}
		item.Name = name
		item.Quantity = quantity
		return item, true, nil // Neues Item erstellt
	} else if err != nil {
		return item, false, err
	}

	// Wenn existiert: Menge erhöhen
	newQuantity := item.Quantity + quantity
	_, err = DB.Exec("UPDATE items SET quantity = $1 WHERE id = $2", newQuantity, item.ID)
	if err != nil {
		return item, false, err
	}

	item.Name = name
	item.Quantity = newQuantity
	return item, false, nil // Item wurde aktualisiert
}