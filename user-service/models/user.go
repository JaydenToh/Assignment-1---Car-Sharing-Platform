package models

type User struct {
	ID             string `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	MembershipTier string `json:"membership_tier"` // e.g., Basic, Premium, VIP
}
