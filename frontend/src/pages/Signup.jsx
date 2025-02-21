// src/components/Signup.jsx
import React, { useState } from "react";
import { registerUser } from "../utils/api";
import "./Signup.css";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";

function Signup() {
  const navigate = useNavigate();
  const [form, setForm] = useState({
    id: "",
    firstName: "",
    lastName: "",
    email: "",
    password: "",
  });

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await registerUser(
        form.id,
        form.firstName,
        form.lastName,
        form.email,
        form.password
      );
      alert(`User registered: ${response.message}`);
      navigate("/login");
    } catch (error) {
      console.error("Error registering user:", error);
      alert(error.message);
    }
  };

  return (
    <div className="signup-container">
      <div className="signup-form-container">
        <h2>Sign Up</h2>
        <p>Create your account in a few easy steps</p>
        <form onSubmit={handleSubmit} className="signup-form">
          <div className="input-group">
            <input
              name="id"
              type="text"
              placeholder="User ID"
              value={form.id}
              onChange={handleChange}
              required
            />
          </div>
          <div className="input-group">
            <input
              name="firstName"
              type="text"
              placeholder="First Name"
              value={form.firstName}
              onChange={handleChange}
              required
            />
          </div>
          <div className="input-group">
            <input
              name="lastName"
              type="text"
              placeholder="Last Name"
              value={form.lastName}
              onChange={handleChange}
              required
            />
          </div>
          <div className="input-group">
            <input
              name="email"
              type="email"
              placeholder="Email Address"
              value={form.email}
              onChange={handleChange}
              required
            />
          </div>
          <div className="input-group">
            <input
              name="password"
              type="password"
              placeholder="Password"
              value={form.password}
              onChange={handleChange}
              required
            />
          </div>

          <button type="submit" className="signup-button">
            Sign Up
          </button>
        </form>

        <div className="login-link">
          Already have an account? <Link to="/login">Log in</Link>
        </div>
      </div>
    </div>
  );
}

export default Signup;
