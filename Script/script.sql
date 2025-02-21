CREATE TABLE my_db.users (
    ID VARCHAR(5) NOT NULL PRIMARY KEY,
    FirstName VARCHAR(30),
    LastName VARCHAR(30),
    Email VARCHAR(50) UNIQUE,
    Password VARCHAR(255),
    MembershipTier ENUM('Basic', 'Premium', 'VIP') DEFAULT 'Basic'
);


CREATE TABLE my_db.vehicles (
    ID VARCHAR(5) NOT NULL PRIMARY KEY,
    Model VARCHAR(50),
    Status ENUM('available', 'booked') DEFAULT 'available'
);

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

CREATE TABLE my_db.billing (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    MembershipTier ENUM('Basic', 'Premium', 'VIP') NOT NULL,
    HourlyRate DECIMAL(10, 2) NOT NULL,
    DiscountPercentage DECIMAL(5, 2) NOT NULL
);

CREATE TABLE my_db.promotions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    min_amount DECIMAL(10, 2) NOT NULL,
    discount_percentage DECIMAL(5, 2) NOT NULL
);

INSERT INTO my_db.users (ID, FirstName, LastName, Email, Password, MembershipTier)
VALUES
    ('01', 'John', 'Tan', 'john@gmail.com', 'drivingcar', 'Basic'),
    ('02', 'Alice', 'Wonder', 'alice@gmail.com', 'flyingman', 'Basic'),
    ('03', 'Bob', 'Morse', 'bob@gmail.com', 'ridingacar', 'Premium'),
    ('04', 'Jane', 'Doe', 'jane@gmail.com', 'drivingwoman', 'VIP'),
    ('05', 'Jayden', 'Toh', 'jayden@gmail.com', 'jaypassword', 'VIP');

INSERT INTO my_db.vehicles (ID, Model, Status)
VALUES
    ('01', 'Tesla', 'available'),
    ('02', 'Nissan', 'booked'),
    ('03', 'Toyota', 'available'),
    ('04', 'BMW', 'booked'),
    ('05', 'Tesla Model S', 'available');

INSERT INTO my_db.billing (MembershipTier, HourlyRate, DiscountPercentage)
VALUES
    ('Basic', 10.00, 0.00),
    ('Premium', 15.00, 10.00),
    ('VIP', 20.00, 20.00);

UPDATE Billing
SET MembershipTier = 'Basic', HourlyRate = 10.00, DiscountPercentage = 0
WHERE ID = 1;

UPDATE Billing
SET MembershipTier = 'Premium', HourlyRate = 8.00, DiscountPercentage = 10
WHERE ID = 2;

UPDATE Billing
SET MembershipTier = 'VIP', HourlyRate = 5.00, DiscountPercentage = 20
WHERE ID = 3;

INSERT INTO my_db.promotions (min_amount, discount_percentage) VALUES (200, 30.00);
INSERT INTO my_db.promotions (min_amount, discount_percentage) VALUES (100, 20.00);
INSERT INTO my_db.promotions (min_amount, discount_percentage) VALUES (50, 10.00);

-- Ensure the foreign key relationships are established correctly
ALTER TABLE my_db.reservations
ADD CONSTRAINT fk_reservation_user
    FOREIGN KEY (UserID) REFERENCES my_db.users(ID) ON DELETE CASCADE,
ADD CONSTRAINT fk_reservation_vehicle
    FOREIGN KEY (VehicleID) REFERENCES my_db.vehicles(ID) ON DELETE CASCADE;

-- Ensure all ENUM values are consistent and valid
ALTER TABLE my_db.users MODIFY COLUMN MembershipTier ENUM('Basic', 'Premium', 'VIP') DEFAULT 'Basic';
ALTER TABLE my_db.billing MODIFY COLUMN MembershipTier ENUM('Basic', 'Premium', 'VIP') NOT NULL;
ALTER TABLE my_db.vehicles MODIFY COLUMN Status ENUM('available', 'booked') DEFAULT 'available';

