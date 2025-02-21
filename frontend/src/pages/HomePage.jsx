import React from "react";
import { Link } from "react-router-dom";
import "./HomePage.css";

function HomePage() {
  return (
    <div className="home-page container">
      <h1>Welcome to the Microservices Demo</h1>
      <p>Select a section from the boxes below.</p>

      <div className="feature-boxes">
        <Link to="/user" className="feature-box user-box">
          <h2>User Management</h2>
          <p>Manage user accounts, profiles, membership status, and more.</p>
        </Link>

        <Link to="/billing" className="feature-box billing-box">
          <h2>Billing Management</h2>
          <p>Calculate, estimate, and generate invoices for rentals.</p>
        </Link>

        <Link to="/vehicles" className="feature-box vehicle-box">
          <h2>Vehicle Management</h2>
          <p>View, book, modify, and cancel vehicle reservations.</p>
        </Link>
      </div>
    </div>
  );
}

export default HomePage;
