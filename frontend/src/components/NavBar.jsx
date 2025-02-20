import React from "react";
import { Link } from "react-router-dom";
import "./NavBar.css"; // optional styling if you create NavBar.css

function NavBar() {
  return (
    <nav className="navbar">
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/vehicles">Vehicle Management</Link>
        </li>
        <li>
          <Link to="/billing">Billing Management</Link>
        </li>
        <li>
          <Link to="/user">Profile</Link>
        </li>
      </ul>
    </nav>
  );
}

export default NavBar;
