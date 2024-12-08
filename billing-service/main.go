package main

import (
	"log"
	"net/http"

	"billing-service/database"
	"billing-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	database.InitDB()
	defer database.CloseDB()

	// Initialize handlers
	handlers.InitBillingHandler(database.DB)

	// Set up routes
	router := mux.NewRouter()
	router.HandleFunc("/calculate-billing", handlers.CalculateBilling).Methods("POST")
	router.HandleFunc("/estimate-billing", handlers.EstimateBilling).Methods("POST")
	router.HandleFunc("/generate-invoice", handlers.GenerateInvoice).Methods("POST")

	// Start the HTTP server
	log.Println("Billing Service running on port 8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}
