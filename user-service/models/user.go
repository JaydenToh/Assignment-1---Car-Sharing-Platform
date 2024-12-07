package models

// User represents the structure of a user in the application
type User struct {
	ID        string `json:"id"`         // User ID (Primary Key)
	FirstName string `json:"first_name"` // User's first name
	LastName  string `json:"last_name"`  // User's last name
	Email     string `json:"email"`      // User's email address (Unique)
	Password  string `json:"password"`   // Hashed password for authentication
}
