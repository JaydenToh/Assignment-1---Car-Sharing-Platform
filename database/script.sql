CREATE TABLE Users (
    ID VARCHAR(5) NOT NULL PRIMARY KEY,
    FirstName VARCHAR(30),
    LastName VARCHAR(30),
    Email VARCHAR(50) UNIQUE,
    Password VARCHAR(255)
);

CREATE TABLE Vehicles (
    ID VARCHAR(5) NOT NULL PRIMARY KEY,
    Model VARCHAR(50),
    Status ENUM('available', 'booked') DEFAULT 'available'
);

CREATE TABLE Reservations (
    ID VARCHAR(5) NOT NULL PRIMARY KEY,
    UserID VARCHAR(5) NOT NULL,
    VehicleID VARCHAR(5) NOT NULL,
    StartTime DATETIME,
    EndTime DATETIME,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    FOREIGN KEY (VehicleID) REFERENCES Vehicles(ID)
);

CREATE TABLE Billing (
    ID VARCHAR(5) NOT NULL PRIMARY KEY,
    ReservationID VARCHAR(5) NOT NULL,
    Amount DECIMAL(10, 2),
    FOREIGN KEY (ReservationID) REFERENCES Reservations(ID)
);

INSERT INTO Users (ID, FirstName, LastName, Email, Password)
VALUES
('01', 'John', 'Tan', 'john@gmail.com', 'drivingcar'),
('02', 'Alice', 'Wonder', 'alice@gmail.com', 'flyingman'),
('03', 'Bob', 'Morse', 'bob@gmail.com', 'ridingacar');

ALTER TABLE Users ADD MembershipTier ENUM('Basic', 'Premium', 'VIP') DEFAULT 'Basic';

INSERT INTO Vehicles (ID, Model, Status)
VALUES
('01', 'Tesla', 'available'),
('02', 'Nissan', 'booked'),
('03', 'Toyota', 'available'),
('04', 'BMW', 'booked');

INSERT INTO Reservations (ID, UserID, VehicleID, StartTime, EndTime)
VALUES
('01', '01', '01', '2024-12-08 09:00:00', '2024-12-08 12:00:00'),
('02', '02', '02', '2024-12-09 14:00:00', '2024-12-09 16:00:00'),
('03', '03', '04', '2024-12-10 10:00:00', '2024-12-10 15:00:00');

INSERT INTO Billing (ID, ReservationID, Amount)
VALUES
('01', '01', 30.00),
('02', '02', 40.00),
('03', '03', 55.00);
