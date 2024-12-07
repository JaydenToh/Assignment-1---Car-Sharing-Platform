package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func loginUser(w http.ResponseWriter, r *http.Request) {
	var credentials User
	_ = json.NewDecoder(r.Body).Decode(&credentials)

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE email = ?", credentials.Email).Scan(&storedPassword)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(credentials.Password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Login successful!")
}
