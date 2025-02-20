import React, { useEffect, useState } from "react";
import "./VehicleManagement.css";
import {
  fetchVehicles,
  bookVehicle,
  modifyBooking,
  cancelBooking,
} from "../utils/api";

function VehicleManagement() {
  const [vehicles, setVehicles] = useState([]);
  const [loading, setLoading] = useState(true);

  // For modifying a reservation
  const [reservationId, setReservationId] = useState("");
  const [modStartTime, setModStartTime] = useState("");
  const [modEndTime, setModEndTime] = useState("");

  useEffect(() => {
    const getData = async () => {
      try {
        const vehicleData = await fetchVehicles();
        setVehicles(vehicleData);
      } catch (error) {
        console.error("Error fetching vehicles:", error);
        alert("Failed to fetch vehicles");
      } finally {
        setLoading(false);
      }
    };
    getData();
  }, []);

  // Book a vehicle
  const handleBookVehicle = async (vehicleId) => {
    // Hard-coded user ID for example
    const userId = "01";
    const startTime = prompt("Enter start time (YYYY-MM-DD HH:MM:SS):");
    const endTime = prompt("Enter end time (YYYY-MM-DD HH:MM:SS):");

    if (!startTime || !endTime) {
      alert("All fields required");
      return;
    }

    try {
      await bookVehicle(vehicleId, userId, startTime, endTime);
      alert("Vehicle booked successfully!");
      // Optionally refresh the vehicle list
    } catch (error) {
      alert(`Error booking vehicle: ${error.message}`);
    }
  };

  // Modify a reservation
  const handleModifyReservation = async () => {
    if (!reservationId || !modStartTime || !modEndTime) {
      alert("All fields required for modifying a reservation");
      return;
    }

    try {
      await modifyBooking(reservationId, modStartTime, modEndTime);
      alert("Reservation modified successfully!");
    } catch (error) {
      alert(`Error modifying reservation: ${error.message}`);
    }
  };

  // Cancel a reservation
  const handleCancelBooking = async () => {
    const resId = prompt("Enter reservation ID to cancel:");
    if (!resId) {
      return;
    }

    try {
      await cancelBooking(resId);
      alert("Reservation canceled successfully!");
    } catch (error) {
      alert(`Error canceling reservation: ${error.message}`);
    }
  };

  if (loading) {
    return <div>Loading vehicles...</div>;
  }

  return (
    <div className="vehicle-management">
      <h1>Vehicle Management</h1>

      <section>
        <h2>Available Vehicles</h2>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Model</th>
              <th>Status</th>
              <th>Book</th>
            </tr>
          </thead>
          <tbody>
            {vehicles.map((vehicle) => (
              <tr key={vehicle.id}>
                <td>{vehicle.id}</td>
                <td>{vehicle.model}</td>
                <td>{vehicle.status}</td>
                <td>
                  {vehicle.status === "available" ? (
                    <button onClick={() => handleBookVehicle(vehicle.id)}>
                      Book
                    </button>
                  ) : (
                    "Not Available"
                  )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </section>

      <hr />

      <section>
        <h2>Modify a Reservation</h2>
        <div>
          <label>Reservation ID:</label>
          <input
            value={reservationId}
            onChange={(e) => setReservationId(e.target.value)}
          />
        </div>
        <div>
          <label>New Start Time (YYYY-MM-DD HH:MM:SS):</label>
          <input
            value={modStartTime}
            onChange={(e) => setModStartTime(e.target.value)}
          />
        </div>
        <div>
          <label>New End Time (YYYY-MM-DD HH:MM:SS):</label>
          <input
            value={modEndTime}
            onChange={(e) => setModEndTime(e.target.value)}
          />
        </div>
        <button onClick={handleModifyReservation}>Modify</button>
      </section>

      <hr />

      <section>
        <h2>Cancel a Reservation</h2>
        <button onClick={handleCancelBooking}>Cancel Reservation</button>
      </section>
    </div>
  );
}

export default VehicleManagement;
