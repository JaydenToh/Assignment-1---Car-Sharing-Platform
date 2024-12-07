package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func initDB() {
	var err error

	// Configure the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),     // Database user
		os.Getenv("DB_PASSWORD"), // Database password
		os.Getenv("DB_HOST"),     // Database host
		os.Getenv("DB_NAME"),     // Database name
	)

	// Open a connection to the database
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	log.Println("Connected to the database successfully.")
}

func closeDB() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed.")
	}
}
