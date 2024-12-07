package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
		// Initialize the database connection
		initDB()
		defer closeDB()
	
		// Set up the router
		router := mux.NewRouter()
		router.HandleFunc("/register", registerUser).Methods("POST")
		router.HandleFunc("/login", loginUser).Methods("POST")
	
		log.Println("User Service running on port 8000")
		log.Fatal(http.ListenAndServe(":8000", router))
}