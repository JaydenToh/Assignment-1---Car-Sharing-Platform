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

CREATE TABLE billing (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    MembershipTier VARCHAR(50) NOT NULL,
    HourlyRate DECIMAL(10, 2) NOT NULL,
    DiscountPercentage DECIMAL(5, 2) NOT NULL
);

ALTER TABLE Billing
ADD COLUMN MembershipTier ENUM('Basic', 'Premium', 'VIP'),
ADD COLUMN HourlyRate DECIMAL(10, 2),
ADD COLUMN DiscountPercentage DECIMAL(5, 2);

UPDATE Billing
SET MembershipTier = 'Basic', HourlyRate = 10.00, DiscountPercentage = 0
WHERE ID = 1;

UPDATE Billing
SET MembershipTier = 'Premium', HourlyRate = 8.00, DiscountPercentage = 10
WHERE ID = 2;

UPDATE Billing
SET MembershipTier = 'VIP', HourlyRate = 5.00, DiscountPercentage = 20
WHERE ID = 3;

INSERT INTO my_db.billing (ReservationID, Amount, MembershipTier, HourlyRate, DiscountPercentage)
VALUES
    (01, 30.00, 'Basic', 10.00, 0.00),
    (02, 40.00, 'Premium', 8.00, 10.00),
    (03, 55.00, 'VIP', 5.00, 20.00);

ALTER TABLE my_db.reservations MODIFY COLUMN ID VARCHAR(36);

ALTER TABLE my_db.billing
DROP FOREIGN KEY billing_ibfk_1;

ALTER TABLE my_db.billing
ADD CONSTRAINT billing_ibfk_1 FOREIGN KEY (ReservationID) REFERENCES my_db.reservations(ID) ON DELETE CASCADE;

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

INSERT INTO billing (MembershipTier, HourlyRate, DiscountPercentage) VALUES
('Basic', 10.00, 0.00),
('Premium', 15.00, 10.00),
('VIP', 20.00, 20.00);

