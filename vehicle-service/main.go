package main

import (
	"log"
	"net/http"

	"vehicle-service/database"
	"vehicle-service/handlers"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// 1. Initialize DB
	database.InitDB()
	defer database.CloseDB()

	// 2. Init handlers
	handlers.InitReservationHandler(database.DB)

	// 3. Setup router
	router := mux.NewRouter()

	// 4. Register routes
	router.HandleFunc("/get-vehicles", handlers.FetchVehicles).Methods("GET")
	router.HandleFunc("/book-vehicle", handlers.BookVehicle).Methods("POST")
	router.HandleFunc("/modify-booking", handlers.ModifyReservation).Methods("PUT")
	router.HandleFunc("/cancel-booking", handlers.CancelBooking).Methods("DELETE")

	// 5. Optional: Set up CORS
	corsOptions := gorillahandlers.AllowedOrigins([]string{"http://localhost:5173"})
	corsMethods := gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	corsHeaders := gorillahandlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// 6. Start server
	log.Println("Vehicle Service running on port 8002")
	log.Fatal(http.ListenAndServe(":8002", gorillahandlers.CORS(
		corsOptions,
		corsMethods,
		corsHeaders,
	)(router)))
}
