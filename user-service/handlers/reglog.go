package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"user-service/models"

	"golang.org/x/crypto/bcrypt"
)

var dbAuth *sql.DB

func InitAuthHandler(database *sql.DB) {
	dbAuth = database
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing /register request")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if user.ID == "" || user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		log.Printf("Error hashing password: %v", hashErr)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	_, dbErr := dbAuth.Exec(
		"INSERT INTO Users (ID, FirstName, LastName, Email, Password) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.FirstName, user.LastName, user.Email, user.Password,
	)
	if dbErr != nil {
		log.Printf("Error inserting user into database: %v", dbErr)
		http.Error(w, "Failed to register user: "+dbErr.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("User registered successfully!")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully!",
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing /login request")

	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if credentials.Email == "" || credentials.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

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

	if bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentials.Password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	log.Println("Login successful!")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful!",
	})
}
