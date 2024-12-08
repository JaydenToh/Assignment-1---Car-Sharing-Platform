package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

var dbProfile *sql.DB // Shared database reference for profile-related handlers

// InitProfileHandler initializes the database for the profile handlers
func InitProfileHandler(database *sql.DB) {
	dbProfile = database
}

// UpdateProfile allows users to update their profile details
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var userUpdate struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	_, dbErr := dbProfile.Exec(
		"UPDATE Users SET FirstName = ?, LastName = ?, Email = ?, Password = ? WHERE ID = ?",
		userUpdate.FirstName, userUpdate.LastName, userUpdate.Email, userUpdate.Password, userUpdate.ID,
	)
	if dbErr != nil {
		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Profile updated successfully!",
	})
}

// GetMembershipStatus allows users to view their membership tier
func GetMembershipStatus(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	var membershipTier string

	err := dbProfile.QueryRow("SELECT MembershipTier FROM Users WHERE ID = ?", userID).Scan(&membershipTier)
	if err != nil {
		http.Error(w, "Failed to retrieve membership status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"membership_tier": membershipTier,
	})
}

// GetRentalHistory retrieves the user's rental history
func GetRentalHistory(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	rows, err := dbProfile.Query(
		"SELECT VehicleID, StartDate, EndDate FROM Reservations WHERE UserID = ?",
		userID,
	)
	if err != nil {
		http.Error(w, "Failed to retrieve rental history", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var history []map[string]interface{}
	for rows.Next() {
		var vehicleID string
		var startDate, endDate string
		if err := rows.Scan(&vehicleID, &startDate, &endDate); err == nil {
			history = append(history, map[string]interface{}{
				"vehicle_id": vehicleID,
				"start_date": startDate,
				"end_date": endDate,
			})
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(history)
}
