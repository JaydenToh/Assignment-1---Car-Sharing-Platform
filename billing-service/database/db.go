package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

var DB *sql.DB // Shared database connection

// InitDB initializes the database connection
func InitDB() {
	var err error

	
		// Set up the environment variables (replace with actual values or set them in your environment)
		os.Setenv("DB_USER", "root")         // Replace with your DB username
		os.Setenv("DB_PASSWORD", "30G776292t05") // Replace with your DB password
		os.Setenv("DB_HOST", "localhost")       // Replace with your DB host and port
		os.Setenv("DB_NAME", "my_db")         // Replace with your database name	

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Validate environment variables
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbName == "" {
		log.Fatal("Missing required database environment variables")
	}

	// DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	// Open database connection
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Connected to the database successfully.")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed.")
	}
}
