package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Connect
func Connect(username string, password string, dbname string, sslmode bool) {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%t", username, password, dbname, sslmode))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}

// Select
func Select(query string) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Update
func Update(query string) (sql.Result, error) {
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Insert
func Insert(query string) (sql.Result, error) {
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Close
func Close() {
	db.Close()
}
