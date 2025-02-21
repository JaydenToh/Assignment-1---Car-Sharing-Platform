# Assignment-1---Car-Sharing-Platform

Creating a Car Sharing Platform in Go

## Overview

This microservice architecture is designed for a **Billing and Payment Processing** system. It handles:

- **Tier-based pricing and discounts**
- **Real-time billing calculation**
- **Secure payment processing**
- **Invoicing and receipts generation**

These services communicate via RESTful APIs, making the system **scalable**, **maintainable**, and **resilient**.

---

## Design Considerations

### Scalability

- Each microservice can be scaled independently.
- Stateless services facilitate horizontal scaling.

### Maintainability

- Separate microservices handle distinct business logic: billing, payments, users, reservations, and vehicles.

### Security

- **HTTPS** used for external communication.
- **API Gateway** or load balancer routes requests securely.
- Sensitive data (payment details, passwords) is encrypted at rest and in transit.

### Fault Tolerance & Resilience

- Microservices are isolated, preventing system-wide failure.

### Technology Stack

- **Backend**: Go (Golang) for microservices
- **Frontend**: React + Vite
- **Database**: MySQL

---

## Architecture Diagram

![Architecture Diagram](./frontend/src/assets/ArchitectureDiagram.png)

- **Billing Service**: Calculates costs, applies discounts, generates invoices.
- **User Service**: Manages user registration, authentication, and membership tiers.
- **Vehicles Service**: Tracks vehicle status (available/booked). Book vehicle, modify booking and cancel reservation
- **MySQL Database**: Stores structured data for each microservice.

---

## Microservice Details

### Billing Service

- **Endpoints**:
  - `POST /calculate-billing`
  - `POST /estimate-billing`
  - `POST /generate-invoice`
- **Usage**: Tier-based pricing, real-time billing calculations, and invoicing.

### Payment Service

- **Endpoints**:
  - `POST /process-payment`
  - `POST /refund`
- **Usage**: Secure transaction handling and refund processing.

### User Service

- **Endpoints**:
  - `POST /register`
  - `POST /login`
  - `GET /user/:id`
- **Usage**: Manages user accounts, membership tiers, and authentication.

### Reservations Service

- **Endpoints**:
  - `POST /create-reservation`
  - `GET /reservations/:userId`
- **Usage**: Oversees booking details, tying users to vehicles with time ranges.

### Vehicles Service

- **Endpoints**:
  - `GET /vehicles`
  - `POST /update-status`
- **Usage**: Tracks vehicle inventory, model details, and availability status.

---

## Database Design

- **`users`**: Holds user info, including name, email, membership tier, etc.
- **`vehicles`**: Stores vehicle data (model, status).
- **`reservations`**: Associates users with vehicles, start/end times, cost details.
- **`billing`**: Contains membership tiers, hourly rates, discount percentages.
- **`promotions`**: Discount percentages.

---

# Setup and Instructions

## Database Setup

### Database Script

A file named **`script.sql`** is provided. This script creates the necessary tables and initial data.

```sql
-- Drop existing tables if they exist
DROP TABLE IF EXISTS my_db.invoices;
DROP TABLE IF EXISTS my_db.billing;
DROP TABLE IF EXISTS my_db.reservations;
DROP TABLE IF EXISTS my_db.vehicles;
DROP TABLE IF EXISTS my_db.users;

-- Create Users Table
CREATE TABLE my_db.users (
    ID VARCHAR(5) NOT NULL PRIMARY KEY,
    FirstName VARCHAR(30),
    LastName VARCHAR(30),
    Email VARCHAR(50) UNIQUE,
    Password VARCHAR(255),
    MembershipTier ENUM('Basic', 'Premium', 'VIP') DEFAULT 'Basic'
);

-- Data for Users
INSERT INTO my_db.users (ID, FirstName, LastName, Email, Password, MembershipTier)
VALUES
    ('01', 'John', 'Tan', 'john@gmail.com', 'drivingcar', 'Basic'),
    ('02', 'Alice', 'Wonder', 'alice@gmail.com', 'flyingman', 'Premium'),
    ('03', 'Bob', 'Morse', 'bob@gmail.com', 'ridingacar', 'Premium'),
    ('04', 'Jane', 'Doe', 'jane@gmail.com', 'drivingwoman', 'VIP'),
    ('05', 'Jayden', 'Toh', 'jayden@gmail.com', 'jaypassword', 'VIP');

-- Create Vehicles Table
CREATE TABLE my_db.vehicles (
    ID VARCHAR(5) NOT NULL PRIMARY KEY,
    Model VARCHAR(50),
    Status ENUM('available', 'booked') DEFAULT 'available'
);

-- Data for Vehicles
INSERT INTO my_db.vehicles (ID, Model, Status)
VALUES
    ('01', 'Tesla', 'available'),
    ('02', 'Nissan', 'booked'),
    ('03', 'Toyota', 'available'),
    ('04', 'BMW', 'booked'),
    ('05', 'Tesla Model S', 'available');

-- Create Reservations Table
CREATE TABLE my_db.reservations (
    ID VARCHAR(36) NOT NULL PRIMARY KEY,
    UserID VARCHAR(5) NOT NULL,
    VehicleID VARCHAR(5) NOT NULL,
    StartTime DATETIME NOT NULL,
    EndTime DATETIME NOT NULL,
    Hours DECIMAL(5, 2) NOT NULL,
    Cost DECIMAL(10, 2) NOT NULL,
    Discount DECIMAL(10, 2) NOT NULL,
    Total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES my_db.users(ID) ON DELETE CASCADE,
    FOREIGN KEY (VehicleID) REFERENCES my_db.vehicles(ID) ON DELETE CASCADE
);

-- Create Promotions Table
CREATE TABLE my_db.promotions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    min_amount DECIMAL(10, 2) NOT NULL,
    discount_percentage DECIMAL(5, 2) NOT NULL
);

-- Data for Promotions
INSERT INTO my_db.promotions (min_amount, discount_percentage) VALUES (200, 30.00);
INSERT INTO my_db.promotions (min_amount, discount_percentage) VALUES (100, 20.00);
INSERT INTO my_db.promotions (min_amount, discount_percentage) VALUES (50, 10.00);

```

## Postman Testing

### 🚀 Billing Service

Base URL: `http://localhost:8001`

---

### 1. Calculate Billing

- **Method:** `POST`
- **Endpoint:** `/calculate-billing`
- **Description:** Calculate the billing amount based on membership tier and rental duration.

**Request Headers:**

```http
Content-Type: application/json

Request Body:
{
  "membership_tier": "Premium",
  "start_time": "2025-03-01 10:00:00",
  "end_time": "2025-03-01 12:00:00"
}

Expected Response (200 OK):
{
  "membership_tier": "Premium",
  "start_time": "2025-03-01 10:00:00",
  "end_time": "2025-03-01 12:00:00",
  "hours": 2,
  "cost": 30.0,
  "discount": 3.0,
  "total": 27.0
}
```
