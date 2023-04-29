package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a database connection pool
	db, err := sql.Open("mysql", "user:password@tcp(db-hostname:3306)/mydb")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Set connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Use the connection pool to perform database operations
	// ...
}
