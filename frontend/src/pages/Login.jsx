// src/components/Login.jsx
import React, { useState } from "react";
import { loginUser } from "../utils/api";
import "./Login.css";

function Login() {
  const [form, setForm] = useState({ email: "", password: "" });

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await loginUser(form.email, form.password);
      alert(`User logged in: ${response.message}`);
    } catch (error) {
      console.error("Error logging in:", error);
      alert(error.message);
    }
  };

  return (
    <div className="login-container">
      <div className="login-form-container">
        <h2>Login</h2>
        <p>Login to your Account Now!</p>
        <form onSubmit={handleSubmit} className="login-form">
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
          <div className="login-options">
            <a href="#" className="forgot-password">
              Forget password?
            </a>
          </div>
          <button type="submit" className="login-button">
            Log in
          </button>
          <p className="signup-link">
            Don’t have an account? <a href="/signup">Sign up</a>
          </p>
        </form>
      </div>
    </div>
  );
}

export default Login;
