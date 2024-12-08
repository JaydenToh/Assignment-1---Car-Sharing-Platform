const BASE_URL_USER = "http://localhost:8000";
const BASE_URL_VEHICLE = "http://localhost:8002";
const BASE_URL_BILLING = "http://localhost:8001";

// User Service API
export const fetchUserProfile = async () => {
  const response = await fetch(`${BASE_URL_USER}/membership-status`);
  if (!response.ok) {
    throw new Error("Failed to fetch user profile");
  }
  return response.json();
};

export const fetchRentalHistory = async () => {
  const response = await fetch(`${BASE_URL_USER}/rental-history`);
  if (!response.ok) {
    throw new Error("Failed to fetch rental history");
  }
  return response.json();
};

// Vehicle Service API
export const fetchReservations = async () => {
  const response = await fetch(`${BASE_URL_VEHICLE}/check-availability`);
  if (!response.ok) {
    throw new Error("Failed to fetch reservations");
  }
  return response.json();
};

export const fetchVehicles = async () => {
  const response = await fetch(`${BASE_URL_VEHICLE}/check-availability`);
  if (!response.ok) {
    throw new Error("Failed to fetch vehicles");
  }
  return response.json();
};

// Billing Service API
export const fetchBilling = async () => {
  const response = await fetch(`${BASE_URL_BILLING}/calculate-billing`);
  if (!response.ok) {
    throw new Error("Failed to fetch billing details");
  }
  return response.json();
};
