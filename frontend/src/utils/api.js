const BASE_URL = "http://localhost:8002"; // point to your Go backend

// 1. Fetch Vehicles (GET /get-vehicles)
export async function fetchVehicles() {
  const response = await fetch(`${BASE_URL}/get-vehicles`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || "Failed to fetch vehicles");
  }
  return await response.json();
}

// 2. Book a Vehicle (POST /book-vehicle)
export async function bookVehicle(vehicleId, userId, startTime, endTime) {
  const response = await fetch(`${BASE_URL}/book-vehicle`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      vehicle_id: vehicleId,
      user_id: userId,
      start_time: startTime,
      end_time: endTime,
    }),
  });
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || "Failed to book vehicle");
  }
  return await response.json();
}

// 3. Modify a Booking (PUT /modify-booking)
export async function modifyBooking(reservationId, newStartTime, newEndTime) {
  const parsedReservationId = parseInt(reservationId, 10); // Ensure it's a number

  const response = await fetch(`${BASE_URL}/modify-booking`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      reservation_id: parsedReservationId, // This should be a number
      start_time: newStartTime,
      end_time: newEndTime,
    }),
  });
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || "Failed to modify booking");
  }
  return await response.json();
}

// 4. Cancel a Booking (DELETE /cancel-booking)
export async function cancelBooking(reservationId) {
  const parsedReservationId = parseInt(reservationId, 10); // Ensure it's a number

  const response = await fetch(`${BASE_URL}/cancel-booking`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      reservation_id: parsedReservationId, // This should be a number
    }),
  });
  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(errorText || "Failed to cancel booking");
  }
  return await response.json();
}

// --------------------
// User Service
// --------------------
export async function registerUser(id, firstName, lastName, email, password) {
  const response = await fetch("http://localhost:8000/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      id,
      first_name: firstName,
      last_name: lastName,
      email,
      password,
    }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to register");
  }
  return response.json();
}

export async function loginUser(email, password) {
  const response = await fetch("http://localhost:8000/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to login");
  }
  return response.json();
}

export async function updateProfile(id, firstName, lastName, email, password) {
  const response = await fetch("http://localhost:8000/update-profile", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      id,
      first_name: firstName,
      last_name: lastName,
      email,
      password,
    }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to update profile");
  }
  return response.json();
}

export async function getMembershipStatus(id) {
  const response = await fetch(
    `http://localhost:8000/membership-status?id=${encodeURIComponent(id)}`
  );
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to get membership status");
  }
  return response.json();
}

export async function getRentalHistory(id) {
  const response = await fetch(
    `http://localhost:8000/rental-history?id=${encodeURIComponent(id)}`
  );
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to get rental history");
  }
  return response.json();
}

export async function calculateBilling(data) {
  const response = await fetch("http://localhost:8001/calculate-billing", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to calculate billing");
  }
  return response.json();
}

export async function estimateBilling(data) {
  const response = await fetch("http://localhost:8001/estimate-billing", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to estimate billing");
  }
  return response.json();
}

export async function generateInvoice(data) {
  const response = await fetch("http://localhost:8001/generate-invoice", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to generate invoice");
  }
  return response.json();
}
