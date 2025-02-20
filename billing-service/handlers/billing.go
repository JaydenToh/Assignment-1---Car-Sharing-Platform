package handlers

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"billing-service/models"
)

var dbBilling *sql.DB

// InitBillingHandler initializes the database for the billing handlers
func InitBillingHandler(database *sql.DB) {
	dbBilling = database
}

// CalculateBilling calculates the cost of a rental
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
	rental.Hours = endTime.Sub(startTime).Hours()

	// Retrieve HourlyRate and DiscountPercentage from Billing table
	var hourlyRate, discountPercentage float64
	err = dbBilling.QueryRow(
		"SELECT HourlyRate, DiscountPercentage FROM my_db.billing WHERE MembershipTier = ?",
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

// EstimateBilling calculates estimated cost without saving the result
func EstimateBilling(w http.ResponseWriter, r *http.Request) {
	var rental struct {
		MembershipTier string `json:"membership_tier"`
		StartTime      string `json:"start_time"`
		EndTime        string `json:"end_time"`
	}
	err := json.NewDecoder(r.Body).Decode(&rental)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Parse start and end time
	startTime, err := time.Parse("2006-01-02 15:04:05", rental.StartTime)
	if err != nil {
		log.Printf("Error parsing start time: %v", err)
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		return
	}
	endTime, err := time.Parse("2006-01-02 15:04:05", rental.EndTime)
	if err != nil {
		log.Printf("Error parsing end time: %v", err)
		http.Error(w, "Invalid end time format", http.StatusBadRequest)
		return
	}

	// Calculate duration in hours
	duration := endTime.Sub(startTime).Hours()

	// Retrieve HourlyRate and DiscountPercentage from Billing table
	var hourlyRate, discountPercentage float64
	err = dbBilling.QueryRow(
		"SELECT HourlyRate, DiscountPercentage FROM my_db.billing WHERE MembershipTier = ?",
		rental.MembershipTier,
	).Scan(&hourlyRate, &discountPercentage)
	if err != nil {
		log.Printf("Failed to retrieve billing info: %v", err)
		http.Error(w, "Failed to retrieve billing info", http.StatusInternalServerError)
		return
	}

	// Calculate cost, discount, and total
	cost := hourlyRate * duration
	discount := cost * (discountPercentage / 100)
	total := cost - discount

	// Respond with estimated billing details
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"membership_tier": rental.MembershipTier,
		"start_time":      rental.StartTime,
		"end_time":        rental.EndTime,
		"hours":           duration,
		"cost":            cost,
		"discount":        discount,
		"total":           total,
	})
}


// GenerateInvoice generates and sends an invoice (simulated)
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

	// Send Email
	err = sendEmail(invoiceRequest.UserEmail, "Your Invoice", fmt.Sprintf(
		"Thank you for your rental!\nReservation ID: %d\nTotal Amount: $%.2f",
		invoiceRequest.ReservationID, invoiceRequest.TotalAmount,
	))
	if err != nil {
		http.Error(w, "Failed to send invoice", http.StatusInternalServerError)
		return
	}

	// Return success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Invoice sent successfully!",
	})
}

// sendEmail sends an email using SMTP
func sendEmail(to, subject, body string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Check if all environment variables are set
	if from == "" || password == "" || smtpHost == "" || smtpPort == "" {
		log.Println("Email environment variables are not set")
		return fmt.Errorf("email environment variables not set")
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Create the email message
	msg := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	// TLS configuration for secure connection
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// Establish connection to the SMTP server
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		log.Printf("Failed to connect to SMTP server: %v", err)
		return err
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Printf("Failed to create SMTP client: %v", err)
		return err
	}

	// Authenticate and send the email
	if err = client.Auth(auth); err != nil {
		log.Printf("Failed to authenticate: %v", err)
		return err
	}

	if err = client.Mail(from); err != nil {
		log.Printf("Failed to set sender: %v", err)
		return err
	}

	if err = client.Rcpt(to); err != nil {
		log.Printf("Failed to set recipient: %v", err)
		return err
	}

	w, err := client.Data()
	if err != nil {
		log.Printf("Failed to send email body: %v", err)
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		log.Printf("Failed to write email body: %v", err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Printf("Failed to close email writer: %v", err)
		return err
	}

	err = client.Quit()
	if err != nil {
		log.Printf("Failed to quit SMTP client: %v", err)
		return err
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}
