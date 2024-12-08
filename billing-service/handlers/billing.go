package handlers

import (
	"billing-service/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func sendEmail(toEmail, subject, body string) error {
	
	fmt.Println("=== Simulated Email ===")
	fmt.Printf("To: %s\n", toEmail)
	fmt.Printf("Subject: %s\n", subject)
	fmt.Printf("Body:\n%s\n", body)
	fmt.Println("========================")

	// Simulate success
	return nil
}


// EstimateBilling calculates the estimated cost based on user input
func EstimateBilling(w http.ResponseWriter, r *http.Request) {
	var billingRequest struct {
		MembershipTier string `json:"membership_tier"`
		StartTime      string `json:"start_time"`
		EndTime        string `json:"end_time"`
	}
	err := json.NewDecoder(r.Body).Decode(&billingRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate MembershipTier
	if billingRequest.MembershipTier != "Basic" && billingRequest.MembershipTier != "Premium" && billingRequest.MembershipTier != "VIP" {
		http.Error(w, "Invalid membership tier", http.StatusBadRequest)
		return
	}

	// Parse Start and End time
	startTime, err := time.Parse("2006-01-02T15:04:05", billingRequest.StartTime)
	if err != nil {
		http.Error(w, "Invalid start time format. Use ISO 8601 format", http.StatusBadRequest)
		return
	}
	endTime, err := time.Parse("2006-01-02T15:04:05", billingRequest.EndTime)
	if err != nil {
		http.Error(w, "Invalid end time format. Use ISO 8601 format", http.StatusBadRequest)
		return
	}

	// Calculate duration in hours
	duration := endTime.Sub(startTime).Hours()
	if duration <= 0 {
		http.Error(w, "End time must be after start time", http.StatusBadRequest)
		return
	}

	// Fetch the hourly rate and discount percentage for the given membership tier
	var hourlyRate, discountPercentage float64
	err = dbBilling.QueryRow(
		"SELECT HourlyRate, DiscountPercentage FROM my_db.billing WHERE MembershipTier = ?",
		billingRequest.MembershipTier,
	).Scan(&hourlyRate, &discountPercentage)
	if err != nil {
		http.Error(w, "Failed to fetch billing rates", http.StatusInternalServerError)
		return
	}

	// Calculate cost and discount
	cost := hourlyRate * duration
	discount := cost * (discountPercentage / 100)
	total := cost - discount

	// Return the estimated billing
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"membership_tier": billingRequest.MembershipTier,
		"start_time":      billingRequest.StartTime,
		"end_time":        billingRequest.EndTime,
		"hours":           duration,
		"cost":            cost,
		"discount":        discount,
		"total":           total,
	})
}

// GenerateInvoice generates an invoice and sends it via email
func GenerateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoiceRequest struct {
		UserEmail      string `json:"user_email"`
		UserID         string `json:"user_id"`
		ReservationID  string `json:"reservation_id"`
		Amount         string `json:"amount"`
	}
	err := json.NewDecoder(r.Body).Decode(&invoiceRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate email and amount
	if invoiceRequest.UserEmail == "" || invoiceRequest.Amount == "" {
		http.Error(w, "User email and amount are required", http.StatusBadRequest)
		return
	}

	// Create the invoice details
	invoiceDetails := fmt.Sprintf(
		`Invoice:
---------------------
Reservation ID: %s
User ID: %s
Total Amount: $%s
---------------------
Thank you for using our service!`,
		invoiceRequest.ReservationID, invoiceRequest.UserID, invoiceRequest.Amount,
	)

	// Send the invoice via email using MailHog
	err = sendEmail(invoiceRequest.UserEmail, "Your Invoice", invoiceDetails)
	if err != nil {
		http.Error(w, "Failed to send invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Invoice sent successfully!",
	})
}
