package models

// Billing represents the billing details in the system
type Billing struct {
	ID            string  `json:"id"`            // Billing ID (Primary Key)
	ReservationID string  `json:"reservation_id"`// Reservation ID
	Amount        float64 `json:"amount"`        // Billing Amount
}
