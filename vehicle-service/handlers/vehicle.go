package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"vehicle-service/models" // Replace with your actual module path
)

var dbVehicle *sql.DB // Shared database reference for vehicle handlers

// InitVehicleHandler initializes the database for the vehicle handlers
func InitVehicleHandler(database *sql.DB) {
	dbVehicle = database
}

// GetAvailableVehicles retrieves all available vehicles
func GetAvailableVehicles(w http.ResponseWriter, r *http.Request) {
	rows, err := dbVehicle.Query("SELECT ID, Model, Status FROM Vehicles WHERE Status = 'available'")
	if err != nil {
		http.Error(w, "Failed to fetch vehicles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var vehicles []models.Vehicle
	for rows.Next() {
		var vehicle models.Vehicle
		rows.Scan(&vehicle.ID, &vehicle.Model, &vehicle.Status)
		vehicles = append(vehicles, vehicle)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

// ReserveVehicle sets a vehicle's status to 'booked'
func ReserveVehicle(w http.ResponseWriter, r *http.Request) {
	vehicleID := r.URL.Query().Get("id")
	if vehicleID == "" {
		http.Error(w, "Vehicle ID is required", http.StatusBadRequest)
		return
	}

	result, err := dbVehicle.Exec("UPDATE Vehicles SET Status = 'booked' WHERE ID = ? AND Status = 'available'", vehicleID)
	if err != nil {
		http.Error(w, "Failed to reserve vehicle", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Vehicle not available or already booked", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Vehicle reserved successfully!",
	})
}
