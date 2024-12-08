const BASE_URL_VEHICLE = "http://localhost:8002";

// Helper function to format datetime into "YYYY-MM-DD HH:mm:ss"
const formatDateTime = (date) => {
  const pad = (n) => (n < 10 ? `0${n}` : n);
  return `${date.getFullYear()}-${pad(date.getMonth() + 1)}-${pad(
    date.getDate()
  )} ${pad(date.getHours())}:${pad(date.getMinutes())}:${pad(
    date.getSeconds()
  )}`;
};

// Fetch all vehicles
export const fetchVehicles = async () => {
  try {
    const response = await fetch(`${BASE_URL_VEHICLE}/get-vehicles`, {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    });
    if (!response.ok) {
      throw new Error(`Failed to fetch vehicles: ${response.statusText}`);
    }
    return response.json();
  } catch (error) {
    console.error("Error fetching vehicles:", error);
    throw error;
  }
};

// Book a vehicle
export const bookVehicle = async (vehicleId, userId, startTime, endTime) => {
  try {
    const response = await fetch(`${BASE_URL_VEHICLE}/book-vehicle`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        vehicle_id: vehicleId,
        user_id: userId,
        start_time: formatDateTime(new Date(startTime)),
        end_time: formatDateTime(new Date(endTime)),
      }),
    });
    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(
        `Failed to book vehicle: ${response.statusText} - ${errorText}`
      );
    }
    return response.json();
  } catch (error) {
    console.error("Error booking vehicle:", error);
    throw error;
  }
};

// Modify a booking
export const modifyBooking = async (reservationId, startTime, endTime) => {
  try {
    const response = await fetch(`${BASE_URL_VEHICLE}/modify-booking`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        reservation_id: reservationId,
        start_time: formatDateTime(new Date(startTime)),
        end_time: formatDateTime(new Date(endTime)),
      }),
    });
    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(
        `Failed to modify booking: ${response.statusText} - ${errorText}`
      );
    }
    return response.json();
  } catch (error) {
    console.error("Error modifying booking:", error);
    throw error;
  }
};

// Cancel a booking
export const cancelBooking = async (reservationId) => {
  if (!reservationId || reservationId.trim() === "") {
    throw new Error("Reservation ID is required to cancel a booking.");
  }

  try {
    const response = await fetch(`${BASE_URL_VEHICLE}/cancel-booking`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        reservation_id: reservationId,
      }),
    });

    if (!response.ok) {
      // Attempt to retrieve additional error details from the response body
      const errorText = await response.text();
      throw new Error(
        `Failed to cancel booking: ${response.statusText} - ${errorText}`
      );
    }

    // Return JSON response
    return await response.json();
  } catch (error) {
    console.error("Error canceling booking:", error);
    throw error;
  }
};
