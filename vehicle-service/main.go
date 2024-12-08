package main

import (
	"log"
	"net/http"
	"vehicle-service/database"
	"vehicle-service/handlers"

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

	// Start the HTTP server
	log.Println("Vehicle Service is running on port 8002")
	log.Fatal(http.ListenAndServe(":8002", router))
}
