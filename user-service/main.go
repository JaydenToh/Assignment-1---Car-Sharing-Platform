package main

import (
	"log"
	"net/http"

	"user-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	initDB()
	defer closeDB()

	// Initialize the authentication handlers with the database connection
	handlers.InitAuthHandler(db)

	// Set up the router
	router := mux.NewRouter()

	// Define routes for Register and Login
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Start the server
	log.Println("User Service running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}