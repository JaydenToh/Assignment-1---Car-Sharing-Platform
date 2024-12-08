package main

import (
	"log"
	"net/http"

	"user-service/database"
	"user-service/handlers"

	gorillahandlers "github.com/gorilla/handlers" // Import gorilla/handlers for CORS
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	database.InitDB()
	defer database.CloseDB()

	// Initialize the authentication handlers with the database connection
	handlers.InitAuthHandler(database.DB)
	handlers.InitProfileHandler(database.DB)

	// Set up the router
	router := mux.NewRouter()

	// Define routes for Register and Login
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Profile management endpoints
	router.HandleFunc("/update-profile", handlers.UpdateProfile).Methods("PUT")
	router.HandleFunc("/membership-status", handlers.GetMembershipStatus).Methods("GET")
	router.HandleFunc("/rental-history", handlers.GetRentalHistory).Methods("GET")

	// Set up CORS middleware using gorillahandlers
	corsHandler := gorillahandlers.CORS(
		gorillahandlers.AllowedOrigins([]string{"http://localhost:5173"}), // Adjust based on your frontend
		gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gorillahandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Start the server with CORS middleware applied
	log.Println("User Service running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", corsHandler(router)))
}
