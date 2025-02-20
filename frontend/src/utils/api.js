// src/utils/api.js

// --------------------
// Vehicle Service
// --------------------
export async function fetchVehicles() {
  const response = await fetch("http://localhost:8002/get-vehicles");
  if (!response.ok) {
    throw new Error("Failed to fetch vehicles");
  }
  return response.json();
}

export async function bookVehicle(vehicleId, userId, startTime, endTime) {
  const response = await fetch("http://localhost:8002/book-vehicle", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ vehicleId, userId, startTime, endTime }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to book vehicle");
  }
  return response.json();
}

export async function modifyBooking(reservationId, newStartTime, newEndTime) {
  const response = await fetch("http://localhost:8002/modify-booking", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ reservationId, newStartTime, newEndTime }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to modify booking");
  }
  return response.json();
}

export async function cancelBooking(reservationId) {
  const response = await fetch("http://localhost:8002/cancel-booking", {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ reservationId }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to cancel booking");
  }
  return response.json();
}

// --------------------
// User Service
// --------------------
export async function registerUser(email, password) {
  const response = await fetch("http://localhost:8000/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
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

export async function updateProfile(name, phone) {
  const response = await fetch("http://localhost:8000/update-profile", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name, phone }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to update profile");
  }
  return response.json();
}

export async function getMembershipStatus() {
  const response = await fetch("http://localhost:8000/membership-status");
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to get membership status");
  }
  return response.json();
}

export async function getRentalHistory() {
  const response = await fetch("http://localhost:8000/rental-history");
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to get rental history");
  }
  return response.json();
}

// --------------------
// Billing Service
// --------------------
export async function calculateBilling(userId) {
  const response = await fetch("http://localhost:8001/calculate-billing", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ userId }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to calculate billing");
  }
  return response.json();
}

export async function estimateBilling(userId) {
  const response = await fetch("http://localhost:8001/estimate-billing", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ userId }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to estimate billing");
  }
  return response.json();
}

export async function generateInvoice(userId) {
  const response = await fetch("http://localhost:8001/generate-invoice", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ userId }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "Failed to generate invoice");
  }
  return response.json();
}
