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
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("root"),
		os.Getenv("30G776292t05"),
		os.Getenv("localhost"),
		os.Getenv("my_db"),
	)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	log.Println("Connected to the database.")
}

func closeDB() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed.")
	}
}