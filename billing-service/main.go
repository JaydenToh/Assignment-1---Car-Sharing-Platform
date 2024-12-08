package main

import (
	"log"
	"net/http"

	"billing-service/database"
	"billing-service/handlers"

	gorillahandlers "github.com/gorilla/handlers"
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

	// Set up CORS middleware
	corsHandler := gorillahandlers.CORS(
		gorillahandlers.AllowedOrigins([]string{"http://localhost:5173"}), // Adjust based on your frontend
		gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gorillahandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Start the HTTP server with CORS middleware applied
	log.Println("Billing Service running on port 8001")
	log.Fatal(http.ListenAndServe(":8001", corsHandler(router)))
}
