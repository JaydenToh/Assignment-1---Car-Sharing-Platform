package main

import (
	"log"
	"net/http"

	"vehicle-service/database"
	"vehicle-service/handlers"

	gorillahandlers "github.com/gorilla/handlers" // Import CORS handlers package
	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	database.InitDB()
	defer database.CloseDB()

	// Initialize handlers with database
	handlers.InitReservationHandler(database.DB)

	// Create a router and define routes
	router := mux.NewRouter()
	router.HandleFunc("/check-availability", handlers.CheckAvailability).Methods("POST")
	router.HandleFunc("/book-vehicle", handlers.BookVehicle).Methods("POST")
	router.HandleFunc("/modify-booking", handlers.ModifyBooking).Methods("PUT")
	router.HandleFunc("/cancel-booking", handlers.CancelBooking).Methods("DELETE")

	// Define CORS options
	corsOptions := gorillahandlers.AllowedOrigins([]string{"http://localhost:5173"}) // Replace with your frontend URL
	corsMethods := gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	corsHeaders := gorillahandlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Start the HTTP server with CORS
	log.Println("Vehicle Service is running on port 8002")
	log.Fatal(http.ListenAndServe(":8002", gorillahandlers.CORS(corsOptions, corsMethods, corsHeaders)(router)))
}
