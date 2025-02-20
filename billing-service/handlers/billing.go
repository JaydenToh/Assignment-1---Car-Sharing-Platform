package handlers

import (
	"billing-service/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

var dbBilling *sql.DB

// InitBillingHandler initializes the database for the billing handlers
func InitBillingHandler(database *sql.DB) {
	dbBilling = database
}

// sendEmail sends an email using SMTP
func sendEmail(to, subject, body string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Check if all environment variables are set
	if from == "" || password == "" || smtpHost == "" || smtpPort == "" {
		return fmt.Errorf("email environment variables not set")
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}

// CalculateBilling calculates the final cost of a rental
func CalculateBilling(w http.ResponseWriter, r *http.Request) {
	var rental models.Rental
	err := json.NewDecoder(r.Body).Decode(&rental)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Parse start and end time
	startTime, err := time.Parse("2006-01-02 15:04:05", rental.StartTime)
	if err != nil {
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", rental.EndTime)
	if err != nil {
		http.Error(w, "Invalid end time format", http.StatusBadRequest)
		return
	}

	// Calculate duration in hours
	duration := endTime.Sub(startTime).Hours()
	if duration <= 0 {
		http.Error(w, "End time must be after start time", http.StatusBadRequest)
		return
	}
	rental.Hours = duration

	// Retrieve HourlyRate and DiscountPercentage from Billing table (Case Insensitive)
	var hourlyRate, discountPercentage float64
	err = dbBilling.QueryRow(
		"SELECT COALESCE(HourlyRate, 0), COALESCE(DiscountPercentage, 0) FROM my_db.billing WHERE LOWER(MembershipTier) = LOWER(?)",
		rental.MembershipTier,
	).Scan(&hourlyRate, &discountPercentage)
	if err != nil {
		http.Error(w, "Failed to retrieve billing info", http.StatusInternalServerError)
		return
	}

	// Calculate cost, discount, and total
	rental.Cost = hourlyRate * rental.Hours
	rental.Discount = rental.Cost * (discountPercentage / 100)
	rental.Total = rental.Cost - rental.Discount

	// Respond with billing details
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rental)
}

// EstimateBilling provides a cost estimate before confirmation
func EstimateBilling(w http.ResponseWriter, r *http.Request) {
	var rental models.Rental
	err := json.NewDecoder(r.Body).Decode(&rental)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use the same logic as CalculateBilling but without saving the result
	CalculateBilling(w, r)
}

// GenerateInvoice generates and sends an invoice via email
func GenerateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoiceRequest struct {
		UserEmail     string  `json:"user_email"`
		UserID        string  `json:"user_id"`
		ReservationID int     `json:"reservation_id"`
		TotalAmount   float64 `json:"total_amount"`
	}
	err := json.NewDecoder(r.Body).Decode(&invoiceRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Send invoice via email
	err = sendEmail(invoiceRequest.UserEmail, "Your Invoice", fmt.Sprintf(
		"Thank you for your rental!\nReservation ID: %d\nTotal Amount: $%.2f",
		invoiceRequest.ReservationID, invoiceRequest.TotalAmount,
	))
	if err != nil {
		http.Error(w, "Failed to send invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Invoice sent successfully!",
	})
}
