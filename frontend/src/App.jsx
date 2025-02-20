import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import NavBar from "./components/NavBar";
import HomePage from "./pages/HomePage";
import UserManagement from "./pages/UserManagement";
import BillingManagement from "./pages/BillingManagement";
import VehicleManagement from "./pages/VehicleManagement";

function App() {
  return (
    <Router>
      <NavBar />
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/user" element={<UserManagement />} />
        <Route path="/billing" element={<BillingManagement />} />
        <Route path="/vehicles" element={<VehicleManagement />} />
      </Routes>
    </Router>
  );
}

export default App;
