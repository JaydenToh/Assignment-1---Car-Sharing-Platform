// src/components/UserManagement.jsx
import React, { useState } from "react";
import "./UserManagement.css";
import {
  updateProfile,
  getMembershipStatus,
  getRentalHistory,
} from "../utils/api";

function UserManagement() {
  const [profileForm, setProfileForm] = useState({
    id: "",
    firstName: "",
    lastName: "",
    email: "",
    password: "",
  });

  const [queryUserId, setQueryUserId] = useState("");
  const [membershipStatus, setMembershipStatus] = useState("");
  const [rentalHistory, setRentalHistory] = useState([]);
  const [activeSection, setActiveSection] = useState("profile");

  const handleChange = (setForm) => (e) => {
    setForm((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleSubmit = (action, formData, successMessage) => async (e) => {
    e.preventDefault();
    try {
      const response = await action(...Object.values(formData));
      alert(`${successMessage}: ${response.message}`);
    } catch (error) {
      console.error("Error:", error);
      alert(error.message);
    }
  };

  const handleGetMembershipStatus = async () => {
    if (!queryUserId) {
      alert("Please enter a User ID for query");
      return;
    }
    try {
      const response = await getMembershipStatus(queryUserId);
      setMembershipStatus(response.membership_tier || response.status || "");
    } catch (error) {
      console.error("Error getting membership status:", error);
      alert(error.message);
    }
  };

  const handleGetRentalHistory = async () => {
    if (!queryUserId) {
      alert("Please enter a User ID for query");
      return;
    }
    try {
      const response = await getRentalHistory(queryUserId);
      setRentalHistory(response.history || []);
    } catch (error) {
      console.error("Error getting rental history:", error);
      alert(error.message);
    }
  };

  return (
    <div className="user-management">
      <aside className="sidebar">
        <h2>SideBar</h2>
        <ul>
          <li onClick={() => setActiveSection("profile")}>Update Profile</li>
          <li onClick={() => setActiveSection("membership")}>
            Membership & Rental History
          </li>
        </ul>
      </aside>

      <main className="main-content">
        {activeSection === "profile" && (
          <form
            onSubmit={handleSubmit(
              updateProfile,
              profileForm,
              "Profile updated"
            )}
            className="profile-form"
          >
            {["id", "firstName", "lastName", "email", "password"].map(
              (field) => (
                <div key={field} className="form-group">
                  <label>{field}:</label>
                  <input
                    name={field}
                    type={field === "password" ? "password" : "text"}
                    value={profileForm[field]}
                    onChange={handleChange(setProfileForm)}
                  />
                </div>
              )
            )}
            <button type="submit">Update Profile</button>
          </form>
        )}

        {activeSection === "membership" && (
          <div className="membership-section">
            <label>User ID:</label>
            <input
              value={queryUserId}
              onChange={(e) => setQueryUserId(e.target.value)}
              className="input-box"
            />
            <button onClick={handleGetMembershipStatus}>
              Check Membership
            </button>
            {membershipStatus && (
              <p>Your membership status: {membershipStatus}</p>
            )}
            <button onClick={handleGetRentalHistory}>
              View Rental History
            </button>
            {rentalHistory.length > 0 && (
              <ul>
                {rentalHistory.map((rental, index) => (
                  <li key={index}>{JSON.stringify(rental)}</li>
                ))}
              </ul>
            )}
          </div>
        )}
      </main>
    </div>
  );
}

export default UserManagement;
