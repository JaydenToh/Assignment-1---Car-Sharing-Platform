package main

import (
	"log"
	"net/http"

	"vehicle-service/handlers" // Replace with your actual module path

	"github.com/gorilla/mux" // Ensure this is installed: go get github.com/gorilla/mux
)

func main() {
	// Initialize the database connection
	initDB()
	defer closeDB()

	// Initialize the vehicle handlers with the database connection
	handlers.InitVehicleHandler(db)

	// Set up the router
	router := mux.NewRouter()

	// Define routes for Vehicle Service
	router.HandleFunc("/vehicles", handlers.GetAvailableVehicles).Methods("GET")
	router.HandleFunc("/reserve", handlers.ReserveVehicle).Methods("POST")

	// Start the server
	log.Println("Vehicle Service running on port 8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}
