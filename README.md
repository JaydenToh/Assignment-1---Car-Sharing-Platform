# Assignment-1---Car-Sharing-Platform
# Jayden Toh Xuan Ming, S10241868J

Creating a car sharing platform in Go

#1 Design Considerations:
The design system uses a microservice architecture to handle vehicle management efficiently. Each service that I have created plays a part to ensure the scalability and reliability of the system.

- Service Seperation:
  Vehicle Service: Handles vehicle availability, booking, modifications, and cancellations
  User Service: Manages user profiles and membership tiers
  Billing Service: Calculates the membership tiers cost and handles billing-related task

- RESTful APIs:
  To connect the back end together with the front end.

- Script:
  One MySQL query to store the structured data of the tables created in the database.

- Database:
  Services have a seperate database folder db.go so that there will not be conflict errors.

- Error Handling:
  Each services includes error handling with descriptive http status codes so errors can be fixed.

#2 Achitecture Diagram


#3 Setup and Running Instructions
Step 1 - Clone the github repository
Step 2 - Set up the database in MYSQL Workbench
         Create database for each services for users, reservations, vehicle, billing
Step 3 - Run each back end Services:
         - User Service
           cd user-service
           go run main.go
         - Vehicle Service
           cd vehicle-service
           go run main.go
         - Billing Service
           cd billing-service
           go run main.go

Step 4 - Run front end:
         cd frontend
         npm install
         npm run dev

Step 5 - Testing:
         http://localhost:5173

        
