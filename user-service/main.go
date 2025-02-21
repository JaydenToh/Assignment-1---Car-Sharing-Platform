package main

import (
	"log"
	"net/http"

	"user-service/database"
	"user-service/handlers"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	handlers.InitAuthHandler(database.DB)
	handlers.InitProfileHandler(database.DB)

	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	router.HandleFunc("/update-profile", handlers.UpdateProfile).Methods("PUT")
	router.HandleFunc("/membership-status", handlers.GetMembershipStatus).Methods("GET")
	router.HandleFunc("/rental-history", handlers.GetRentalHistory).Methods("GET")

	corsHandler := gorillahandlers.CORS(
		gorillahandlers.AllowedOrigins([]string{"http://localhost:5173"}),
		gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gorillahandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		gorillahandlers.AllowCredentials(),
	)

	log.Println("User Service running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", corsHandler(router)))
}
