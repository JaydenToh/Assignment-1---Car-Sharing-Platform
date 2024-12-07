package handlers

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"billing-service/models" // Replace with your actual module path
)

var dbBilling *sql.DB // Shared database reference for billing handlers

// InitBillingHandler initializes the database for the billing handlers
func InitBillingHandler(database *sql.DB) {
	dbBilling = database
	rand.Seed(time.Now().UnixNano()) // Seed for generating random IDs
}

// GenerateBilling calculates and stores the billing for a reservation
func GenerateBilling(w http.ResponseWriter, r *http.Request) {
	var billing models.Billing
	err := json.NewDecoder(r.Body).Decode(&billing)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a random ID for the billing
	billing.ID = "B" + strconv.Itoa(rand.Intn(10000))

	// Example calculation: Fixed rate * duration (in hours)
	const hourlyRate = 10.0
	var durationInHours int
	err = dbBilling.QueryRow(
		"SELECT TIMESTAMPDIFF(HOUR, StartTime, EndTime) FROM Reservations WHERE ID = ?",
		billing.ReservationID,
	).Scan(&durationInHours)
	if err != nil {
		http.Error(w, "Failed to calculate duration for reservation", http.StatusInternalServerError)
		return
	}
	billing.Amount = float64(durationInHours) * hourlyRate

	// Insert billing details into the database
	_, dbErr := dbBilling.Exec(
		"INSERT INTO Billing (ID, ReservationID, Amount) VALUES (?, ?, ?)",
		billing.ID, billing.ReservationID, billing.Amount,
	)
	if dbErr != nil {
		http.Error(w, "Failed to generate billing: "+dbErr.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the billing details
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(billing)
}

// GetBilling retrieves billing details for a specific reservation
func GetBilling(w http.ResponseWriter, r *http.Request) {
	reservationID := r.URL.Query().Get("reservation_id")
	if reservationID == "" {
		http.Error(w, "Reservation ID is required", http.StatusBadRequest)
		return
	}

	var billing models.Billing
	err := dbBilling.QueryRow(
		"SELECT ID, ReservationID, Amount FROM Billing WHERE ReservationID = ?",
		reservationID,
	).Scan(&billing.ID, &billing.ReservationID, &billing.Amount)
	if err != nil {
		http.Error(w, "Billing details not found", http.StatusNotFound)
		return
	}

	// Respond with the billing details
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(billing)
}
