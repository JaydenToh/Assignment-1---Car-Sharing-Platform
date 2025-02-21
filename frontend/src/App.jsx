import React from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  useLocation,
} from "react-router-dom";

import NavBar from "./components/NavBar";
import HomePage from "./pages/HomePage";
import UserManagement from "./pages/UserManagement";
import BillingManagement from "./pages/BillingManagement";
import VehicleManagement from "./pages/VehicleManagement";
import Signup from "./pages/Signup";
import Login from "./pages/Login";

// Custom Layout Wrapper to control NavBar visibility
const Layout = ({ children }) => {
  const location = useLocation();
  const hideNavRoutes = ["/login", "/signup"];

  // Check if current route is in hideNavRoutes array
  const shouldHideNavBar = hideNavRoutes.includes(location.pathname);

  return (
    <>
      {!shouldHideNavBar && <NavBar />}
      {children}
    </>
  );
};

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/login" element={<Login />} />
          <Route path="/user" element={<UserManagement />} />
          <Route path="/billing" element={<BillingManagement />} />
          <Route path="/vehicles" element={<VehicleManagement />} />
        </Routes>
      </Layout>
    </Router>
  );
}

export default App;
