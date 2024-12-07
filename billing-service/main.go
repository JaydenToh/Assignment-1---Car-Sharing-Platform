package main

import (
	"log"
	"net/http"

	"billing-service/handlers" // Replace with your actual module path

	"github.com/gorilla/mux" // Ensure this is installed: go get github.com/gorilla/mux
)

func main() {
	// Initialize the database connection
	initDB()
	defer closeDB()

	// Initialize the billing handlers with the database connection
	handlers.InitBillingHandler(db)

	// Set up the router
	router := mux.NewRouter()

	// Define routes for Billing Service
	router.HandleFunc("/billing/generate", handlers.GenerateBilling).Methods("POST")
	router.HandleFunc("/billing", handlers.GetBilling).Methods("GET")

	// Start the server
	log.Println("Billing Service running on port 8002")
	log.Fatal(http.ListenAndServe(":8002", router))
}
