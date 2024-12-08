package models

type Invoice struct {
	ID            int     `json:"id"`
	UserID        string  `json:"user_id"`
	ReservationID int     `json:"reservation_id"`
	TotalAmount   float64 `json:"total_amount"`
	CreatedAt     string  `json:"created_at"`
}
