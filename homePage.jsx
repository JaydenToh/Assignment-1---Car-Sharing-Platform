import React, { useEffect, useState } from "react";
import { fetchReservations, fetchVehicles, fetchBilling } from "./utils/api";

function HomePage() {
  const [reservations, setReservations] = useState([]);
  const [vehicles, setVehicles] = useState([]);
  const [billing, setBilling] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [resResponse, vehResponse, billResponse] = await Promise.all([
          fetchReservations(),
          fetchVehicles(),
          fetchBilling(),
        ]);

        setReservations(resResponse);
        setVehicles(vehResponse);
        setBilling(billResponse);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) {
    return <div>Loading data...</div>;
  }

  return (
    <div>
      <h1>Car Sharing Platform - Dashboard</h1>

      {/* Reservations Section */}
      <section>
        <h2>Reservations</h2>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>User ID</th>
              <th>Vehicle ID</th>
              <th>Start Time</th>
              <th>End Time</th>
            </tr>
          </thead>
          <tbody>
            {reservations.map((res) => (
              <tr key={res.id}>
                <td>{res.id}</td>
                <td>{res.userId}</td>
                <td>{res.vehicleId}</td>
                <td>{res.startTime}</td>
                <td>{res.endTime}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </section>

      {/* Vehicles Section */}
      <section>
        <h2>Vehicles</h2>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Model</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {vehicles.map((veh) => (
              <tr key={veh.id}>
                <td>{veh.id}</td>
                <td>{veh.model}</td>
                <td>{veh.status}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </section>

      {/* Billing Section */}
      <section>
        <h2>Billing Details</h2>
        <table>
          <thead>
            <tr>
              <th>Reservation ID</th>
              <th>Amount</th>
              <th>Membership Tier</th>
              <th>Hourly Rate</th>
              <th>Discount</th>
            </tr>
          </thead>
          <tbody>
            {billing.map((bill) => (
              <tr key={bill.id}>
                <td>{bill.reservationId}</td>
                <td>{bill.amount}</td>
                <td>{bill.membershipTier}</td>
                <td>{bill.hourlyRate}</td>
                <td>{bill.discountPercentage}%</td>
              </tr>
            ))}
          </tbody>
        </table>
      </section>
    </div>
  );
}

export default HomePage;
