package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set up the environment variables (replace with actual values or set them in your environment)
	os.Setenv("DB_USER", "root")         // Replace with your DB username
	os.Setenv("DB_PASSWORD", "30G776292t05") // Replace with your DB password
	os.Setenv("DB_HOST", "localhost")       // Replace with your DB host and port
	os.Setenv("DB_NAME", "my_db")         // Replace with your database name

	// Configure the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),     // Database user
		os.Getenv("DB_PASSWORD"), // Database password
		os.Getenv("DB_HOST"),     // Database host
		os.Getenv("DB_NAME"),     // Database name
	)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Success message
	log.Println("Connected to the database successfully!")

	// Close the database connection
	defer db.Close()
}
