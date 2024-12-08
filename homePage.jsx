import React, { useEffect, useState } from "react";
import {
  fetchVehicles,
  bookVehicle,
  modifyBooking,
  cancelBooking,
} from "./utils/api";
import "./styles.css";

function VehicleManagement() {
  const [vehicles, setVehicles] = useState([]);
  const [loading, setLoading] = useState(true);

  // Fetch vehicles on component mount
  useEffect(() => {
    const fetchData = async () => {
      try {
        const vehicleData = await fetchVehicles();
        setVehicles(vehicleData);
      } catch (error) {
        console.error("Error fetching vehicles:", error);
        alert("Failed to fetch vehicles. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  // Book a vehicle
  const handleBookVehicle = async () => {
    const vehicleId = prompt("Enter vehicle ID to book:");
    const userId = prompt("Enter your User ID:");
    const startTime = prompt("Enter start time (YYYY-MM-DD HH:MM:SS):");
    const endTime = prompt("Enter end time (YYYY-MM-DD HH:MM:SS):");

    if (!vehicleId || !userId || !startTime || !endTime) {
      alert("All fields are required to book a vehicle.");
      return;
    }

    try {
      const response = await bookVehicle(vehicleId, userId, startTime, endTime);
      alert(`Vehicle booked successfully: ${response.message}`);
    } catch (error) {
      console.error("Error booking vehicle:", error);
      alert(`Error booking vehicle: ${error.message}`);
    }
  };

  // Modify a booking
  const handleModifyBooking = async () => {
    const reservationId = prompt("Enter reservation ID to modify:");
    const newStartTime = prompt("Enter new start time (YYYY-MM-DD HH:MM:SS):");
    const newEndTime = prompt("Enter new end time (YYYY-MM-DD HH:MM:SS):");

    if (!reservationId || !newStartTime || !newEndTime) {
      alert("All fields are required to modify a booking.");
      return;
    }

    try {
      const response = await modifyBooking(
        reservationId,
        newStartTime,
        newEndTime
      );
      alert(`Booking modified successfully: ${response.message}`);
    } catch (error) {
      console.error("Error modifying booking:", error);
      alert(`Error modifying booking: ${error.message}`);
    }
  };

  // Cancel a booking
  const handleCancelBooking = async () => {
    const reservationId = prompt("Enter reservation ID to cancel:");

    if (!reservationId) {
      alert("Reservation ID is required to cancel a booking.");
      return;
    }

    try {
      const response = await cancelBooking(reservationId);
      alert(`Booking canceled successfully: ${response.message}`);
    } catch (error) {
      console.error("Error canceling booking:", error);
      alert(`Error canceling booking: ${error.message}`);
    }
  };

  if (loading) {
    return <div>Loading vehicles...</div>;
  }

  return (
    <div className="container">
      <h1>Vehicle Management</h1>
      <section>
        <h2>Available Vehicles</h2>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Model</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {vehicles.map((vehicle) => (
              <tr key={vehicle.id}>
                <td>{vehicle.id}</td>
                <td>{vehicle.model}</td>
                <td>{vehicle.status}</td>
              </tr>
            ))}
          </tbody>
        </table>
        <div className="buttons">
          <button onClick={handleBookVehicle}>Book Vehicle</button>
          <button onClick={handleModifyBooking}>Modify Booking</button>
          <button onClick={handleCancelBooking}>Cancel Booking</button>
        </div>
      </section>
    </div>
  );
}

export default VehicleManagement;
