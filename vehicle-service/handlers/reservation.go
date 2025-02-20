package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var dbReservation *sql.DB

// InitReservationHandler initializes the database handle for reservation handlers
func InitReservationHandler(database *sql.DB) {
	dbReservation = database
}

// 1. FetchVehicles (GET /get-vehicles)
func FetchVehicles(w http.ResponseWriter, r *http.Request) {
	// Query vehicles table for available vehicles
	rows, err := dbReservation.Query(`
		SELECT ID, Model, Status
		FROM vehicles
		WHERE Status = 'available'
	`)
	if err != nil {
		log.Printf("Error in FetchVehicles query: %v", err)
		http.Error(w, "Failed to fetch vehicles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var vehicles []map[string]interface{}

	for rows.Next() {
		var id, model, status string
		// If ID in your table is INT, you can scan into an int. But scanning into string usually works fine for numeric columns in MySQL.
		if err := rows.Scan(&id, &model, &status); err != nil {
			log.Printf("Error scanning vehicles row: %v", err)
			http.Error(w, "Error scanning vehicles", http.StatusInternalServerError)
			return
		}

		vehicles = append(vehicles, map[string]interface{}{
			"id":     id,
			"model":  model,
			"status": status,
		})
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(vehicles); err != nil {
		log.Printf("Error encoding vehicles response: %v", err)
	}
}

// 2. BookVehicle (POST /book-vehicle)
func BookVehicle(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID    string `json:"user_id"`
		VehicleID string `json:"vehicle_id"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Insert into reservations
	result, err := dbReservation.Exec(`
		INSERT INTO reservations (UserID, VehicleID, StartTime, EndTime)
		VALUES (?, ?, ?, ?)
	`, req.UserID, req.VehicleID, req.StartTime, req.EndTime)
	if err != nil {
		log.Printf("Failed to create reservation: %v", err)
		http.Error(w, "Failed to book vehicle", http.StatusInternalServerError)
		return
	}

	newResID, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to get new ResID: %v", err)
		http.Error(w, "Failed to book vehicle (ID retrieval)", http.StatusInternalServerError)
		return
	}

	// Update vehicle status to 'booked'
	_, err = dbReservation.Exec(`
		UPDATE vehicles
		SET Status = 'booked'
		WHERE ID = ?
	`, req.VehicleID)
	if err != nil {
		log.Printf("Failed to update vehicle status: %v", err)
		http.Error(w, "Failed to update vehicle status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "Vehicle booked successfully!",
		"reservationId": newResID,
	})
}

// 3. ModifyReservation (PUT /modify-booking)
func ModifyReservation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ReservationID int    `json:"reservation_id"`
		StartTime     string `json:"start_time"`
		EndTime       string `json:"end_time"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.ReservationID == 0 {
		http.Error(w, "Reservation ID is required", http.StatusBadRequest)
		return
	}

	_, err := dbReservation.Exec(`
		UPDATE reservations
		SET StartTime = ?, EndTime = ?
		WHERE ResID = ?
	`, req.StartTime, req.EndTime, req.ReservationID)
	if err != nil {
		log.Printf("Failed to modify reservation: %v", err)
		http.Error(w, "Failed to modify reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Reservation modified successfully",
	})
}

// 4. CancelBooking (DELETE /cancel-booking)
func CancelBooking(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ReservationID int `json:"reservation_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.ReservationID == 0 {
		http.Error(w, "Reservation ID is required", http.StatusBadRequest)
		return
	}

	// Find vehicle ID for the given reservation
	var vehicleID string
	err := dbReservation.QueryRow(`
		SELECT VehicleID
		FROM reservations
		WHERE ResID = ?
	`, req.ReservationID).Scan(&vehicleID)
	if err != nil {
		log.Printf("Failed to find reservation: %v", err)
		http.Error(w, "Reservation not found", http.StatusNotFound)
		return
	}

	// Delete reservation
	_, err = dbReservation.Exec(`
		DELETE FROM reservations
		WHERE ResID = ?
	`, req.ReservationID)
	if err != nil {
		log.Printf("Failed to cancel reservation: %v", err)
		http.Error(w, "Failed to cancel reservation", http.StatusInternalServerError)
		return
	}

	// Set the vehicle status back to 'available'
	_, err = dbReservation.Exec(`
		UPDATE vehicles
		SET Status = 'available'
		WHERE ID = ?
	`, vehicleID)
	if err != nil {
		log.Printf("Failed to update vehicle status: %v", err)
		http.Error(w, "Failed to update vehicle status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Reservation canceled successfully!",
	})
}
