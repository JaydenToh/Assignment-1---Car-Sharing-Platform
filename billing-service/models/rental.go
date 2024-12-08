package models

type Rental struct {
	UserID         string  `json:"user_id"`
	MembershipTier string  `json:"membership_tier"`
	VehicleID      string  `json:"vehicle_id"`
	StartTime      string  `json:"start_time"`
	EndTime        string  `json:"end_time"`
	Hours          float64 `json:"hours"`
	Cost           float64 `json:"cost"`
	Discount       float64 `json:"discount"`
	Total          float64 `json:"total"`
}
