package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"user-service/models" // Replace with your actual module path

	"golang.org/x/crypto/bcrypt" // Ensure this package is installed via `go get golang.org/x/crypto/bcrypt`
)

var dbAuth *sql.DB // Shared database reference for authentication handlers

// InitAuthHandler initializes the database for the authentication handlers
func InitAuthHandler(database *sql.DB) {
	dbAuth = database
}

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing /register request")

	// Parse incoming JSON payload into the User struct
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if user.ID == "" || user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Hash the password for security
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		log.Printf("Error hashing password: %v", hashErr)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Insert the user into the database
	_, dbErr := dbAuth.Exec(
		"INSERT INTO Users (ID, FirstName, LastName, Email, Password) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.FirstName, user.LastName, user.Email, user.Password,
	)
	if dbErr != nil {
		log.Printf("Error inserting user into database: %v", dbErr)
		http.Error(w, "Failed to register user: "+dbErr.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	log.Println("User registered successfully!")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully!",
	})
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing /login request")

	// Parse incoming JSON payload into the User struct
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if credentials.Email == "" || credentials.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Query for the stored hashed password based on the provided email
	var storedPassword string
	queryErr := dbAuth.QueryRow(
		"SELECT Password FROM Users WHERE Email = ?", credentials.Email,
	).Scan(&storedPassword)
	if queryErr == sql.ErrNoRows {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else if queryErr != nil {
		log.Printf("Error querying database: %v", queryErr)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the stored hashed password
	if bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentials.Password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Respond with a success message
	log.Println("Login successful!")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful!",
	})
}
