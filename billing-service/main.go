package main

import (
	"log"
	"net/http"
	"os"

	"billing-service/database"
	"billing-service/handlers"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	database.InitDB()
	defer database.CloseDB()

	// Initialize billing handlers with the database
	handlers.InitBillingHandler(database.DB)

	// Set up routes without /api/v1 prefix
	router := mux.NewRouter()

	// Billing Routes
	router.HandleFunc("/calculate-billing", handlers.CalculateBilling).Methods("POST")
	router.HandleFunc("/estimate-billing", handlers.EstimateBilling).Methods("POST") // Added Back
	router.HandleFunc("/generate-invoice", handlers.GenerateInvoice).Methods("POST")

	// Health Check Endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Billing Service is running."))
	}).Methods("GET")

	// CORS Configuration
	corsHandler := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:5173"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Start the HTTP server with CORS
	port := os.Getenv("BILLING_PORT")
	if port == "" {
		port = "8001" // Default port if not specified
	}
	log.Printf("Billing Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler(router)))
}
