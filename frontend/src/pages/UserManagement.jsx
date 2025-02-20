import React, { useState } from "react";
import {
  registerUser,
  loginUser,
  updateProfile,
  getMembershipStatus,
  getRentalHistory,
} from "../utils/api";

function UserManagement() {
  const [registerForm, setRegisterForm] = useState({ email: "", password: "" });
  const [loginForm, setLoginForm] = useState({ email: "", password: "" });
  const [profileForm, setProfileForm] = useState({ name: "", phone: "" });
  const [membershipStatus, setMembershipStatus] = useState("");
  const [rentalHistory, setRentalHistory] = useState([]);

  // Handlers for register
  const handleRegisterChange = (e) => {
    setRegisterForm({ ...registerForm, [e.target.name]: e.target.value });
  };
  const handleRegisterSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await registerUser(
        registerForm.email,
        registerForm.password
      );
      alert(`User registered: ${response.message}`);
    } catch (error) {
      console.error("Error registering user:", error);
      alert(error.message);
    }
  };

  // Handlers for login
  const handleLoginChange = (e) => {
    setLoginForm({ ...loginForm, [e.target.name]: e.target.value });
  };
  const handleLoginSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await loginUser(loginForm.email, loginForm.password);
      alert(`User logged in: ${response.message}`);
    } catch (error) {
      console.error("Error logging in user:", error);
      alert(error.message);
    }
  };

  // Handlers for update profile
  const handleProfileChange = (e) => {
    setProfileForm({ ...profileForm, [e.target.name]: e.target.value });
  };
  const handleProfileSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await updateProfile(profileForm.name, profileForm.phone);
      alert(`Profile updated: ${response.message}`);
    } catch (error) {
      console.error("Error updating profile:", error);
      alert(error.message);
    }
  };

  // Handler for membership status
  const handleGetMembershipStatus = async () => {
    try {
      const response = await getMembershipStatus();
      setMembershipStatus(response.status);
    } catch (error) {
      console.error("Error getting membership status:", error);
      alert(error.message);
    }
  };

  // Handler for rental history
  const handleGetRentalHistory = async () => {
    try {
      const response = await getRentalHistory();
      setRentalHistory(response.history || []);
    } catch (error) {
      console.error("Error getting rental history:", error);
      alert(error.message);
    }
  };

  return (
    <div className="container">
      <h1>User Management</h1>

      <section>
        <h2>Register</h2>
        <form onSubmit={handleRegisterSubmit}>
          <div>
            <label>Email:</label>
            <input
              name="email"
              value={registerForm.email}
              onChange={handleRegisterChange}
            />
          </div>
          <div>
            <label>Password:</label>
            <input
              name="password"
              type="password"
              value={registerForm.password}
              onChange={handleRegisterChange}
            />
          </div>
          <button type="submit">Register</button>
        </form>
      </section>

      <section>
        <h2>Login</h2>
        <form onSubmit={handleLoginSubmit}>
          <div>
            <label>Email:</label>
            <input
              name="email"
              value={loginForm.email}
              onChange={handleLoginChange}
            />
          </div>
          <div>
            <label>Password:</label>
            <input
              name="password"
              type="password"
              value={loginForm.password}
              onChange={handleLoginChange}
            />
          </div>
          <button type="submit">Login</button>
        </form>
      </section>

      <section>
        <h2>Update Profile</h2>
        <form onSubmit={handleProfileSubmit}>
          <div>
            <label>Name:</label>
            <input
              name="name"
              value={profileForm.name}
              onChange={handleProfileChange}
            />
          </div>
          <div>
            <label>Phone:</label>
            <input
              name="phone"
              value={profileForm.phone}
              onChange={handleProfileChange}
            />
          </div>
          <button type="submit">Update Profile</button>
        </form>
      </section>

      <section>
        <h2>Membership Status</h2>
        <button onClick={handleGetMembershipStatus}>Check Membership</button>
        {membershipStatus && <p>Your status: {membershipStatus}</p>}
      </section>

      <section>
        <h2>Rental History</h2>
        <button onClick={handleGetRentalHistory}>View Rental History</button>
        {rentalHistory.length > 0 && (
          <ul>
            {rentalHistory.map((rental, index) => (
              <li key={index}>{JSON.stringify(rental)}</li>
            ))}
          </ul>
        )}
      </section>
    </div>
  );
}

export default UserManagement;
