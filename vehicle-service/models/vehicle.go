package models

// Vehicle represents the vehicle in the system
type Vehicle struct {
	ID     string `json:"id"`     // Vehicle ID (Primary Key)
	Model  string `json:"model"`  // Vehicle Model
	Status string `json:"status"` // Vehicle Status (e.g., available or booked)
}
