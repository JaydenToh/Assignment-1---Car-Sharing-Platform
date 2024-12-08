package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var dbReservation *sql.DB // Shared database reference for reservation handlers

// InitReservationHandler initializes the database for reservation handlers
func InitReservationHandler(database *sql.DB) {
	dbReservation = database
}

// CheckAvailability checks vehicle availability
func CheckAvailability(w http.ResponseWriter, r *http.Request) {
	var request struct {
		VehicleID string `json:"vehicle_id"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Parse Start and End times
	startTime, err := time.Parse("2006-01-02 15:04:05", request.StartTime)
	if err != nil {
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", request.EndTime)
	if err != nil {
		http.Error(w, "Invalid end time format", http.StatusBadRequest)
		return
	}

	// Check if the vehicle is available
	var count int
	err = dbReservation.QueryRow(
		`SELECT COUNT(*) FROM reservations 
		 WHERE VehicleID = ? AND 
		 (StartTime < ? AND EndTime > ?)`,
		request.VehicleID, endTime, startTime,
	).Scan(&count)
	if err != nil {
		http.Error(w, "Failed to check availability", http.StatusInternalServerError)
		return
	}

	if count > 0 {
		http.Error(w, "Vehicle is not available for the specified time", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Vehicle is available",
	})
}

func GenerateID() string {
	return uuid.New().String()
}

// BookVehicle handles vehicle booking
func BookVehicle(w http.ResponseWriter, r *http.Request) {
	var booking struct {
		UserID    string `json:"user_id"`
		VehicleID string `json:"vehicle_id"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Parse Start and End time
	startTime, err := time.Parse("2006-01-02 15:04:05", booking.StartTime)
	if err != nil {
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", booking.EndTime)
	if err != nil {
		http.Error(w, "Invalid end time format", http.StatusBadRequest)
		return
	}

	// Check if the vehicle is available
	var status string
	err = dbReservation.QueryRow("SELECT Status FROM vehicles WHERE ID = ?", booking.VehicleID).Scan(&status)
	if err != nil {
		http.Error(w, "Vehicle not found", http.StatusNotFound)
		return
	}
	if status != "available" {
		http.Error(w, "Vehicle is not available", http.StatusConflict)
		return
	}

	// Insert reservation
	_, err = dbReservation.Exec(
		"INSERT INTO reservations (ID, UserID, VehicleID, StartTime, EndTime) VALUES (?, ?, ?, ?, ?)",
		GenerateID(), booking.UserID, booking.VehicleID, startTime, endTime,
	)
	if err != nil {
		log.Printf("Failed to create reservation: %v", err)
		http.Error(w, "Failed to book vehicle", http.StatusInternalServerError)
		return
	}

	// Update vehicle status
	_, err = dbReservation.Exec("UPDATE vehicles SET Status = 'booked' WHERE ID = ?", booking.VehicleID)
	if err != nil {
		log.Printf("Failed to update vehicle status: %v", err)
		http.Error(w, "Failed to update vehicle status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Vehicle booked successfully!",
	})
}

// ModifyBooking modifies an existing reservation
func ModifyBooking(w http.ResponseWriter, r *http.Request) {
	var update struct {
		ReservationID string `json:"reservation_id"`
		StartTime     string `json:"start_time"`
		EndTime       string `json:"end_time"`
	}

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the reservation
	_, err = dbReservation.Exec(
		`UPDATE reservations 
		 SET StartTime = ?, EndTime = ?
		 WHERE ID = ?`,
		update.StartTime, update.EndTime, update.ReservationID,
	)
	if err != nil {
		http.Error(w, "Failed to modify reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Reservation updated successfully",
	})
}

// CancelBooking cancels a reservation
func CancelBooking(w http.ResponseWriter, r *http.Request) {
    // Parse the request body
    var request struct {
        ReservationID string `json:"reservation_id"`
    }
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Retrieve vehicle ID from the reservation
    var vehicleID string
    err = dbReservation.QueryRow(
        "SELECT VehicleID FROM reservations WHERE ID = ?",
        request.ReservationID,
    ).Scan(&vehicleID)
    if err != nil {
        http.Error(w, "Failed to find reservation: "+err.Error(), http.StatusNotFound)
        return
    }

    // Delete the reservation
    _, err = dbReservation.Exec("DELETE FROM reservations WHERE ID = ?", request.ReservationID)
    if err != nil {
        http.Error(w, "Failed to cancel reservation", http.StatusInternalServerError)
        return
    }

    // Update vehicle status to "available"
    _, err = dbReservation.Exec("UPDATE vehicles SET Status = 'available' WHERE ID = ?", vehicleID)
    if err != nil {
        http.Error(w, "Failed to update vehicle status", http.StatusInternalServerError)
        return
    }

    // Success response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Reservation cancelled successfully!",
    })
}