import React, { useState } from "react";
import "./BillingManagement.css";
import {
  calculateBilling,
  estimateBilling,
  generateInvoice,
} from "../utils/api";

function BillingManagement() {
  const [userId, setUserId] = useState("");
  const [membershipTier, setMembershipTier] = useState("Basic");
  const [startTime, setStartTime] = useState("");
  const [endTime, setEndTime] = useState("");
  const [userEmail, setUserEmail] = useState("");
  const [billingResult, setBillingResult] = useState(null);
  const [estimatedCost, setEstimatedCost] = useState(null);
  const [invoiceData, setInvoiceData] = useState(null);

  const formatDateTime = (datetime) => {
    if (!datetime) return "";
    return datetime.replace("T", " ") + ":00";
  };

  // Calculate billing
  const handleCalculateBilling = async () => {
    if (!userId || !membershipTier || !startTime || !endTime) {
      alert("Please fill in all required fields.");
      return;
    }
    try {
      const formattedStartTime = formatDateTime(startTime);
      const formattedEndTime = formatDateTime(endTime);

      const response = await calculateBilling({
        membership_tier: membershipTier,
        start_time: formattedStartTime,
        end_time: formattedEndTime,
      });
      setBillingResult(response);
    } catch (error) {
      console.error("Error calculating billing:", error);
      alert(error.message);
    }
  };

  // Estimate billing
  const handleEstimateBilling = async () => {
    if (!userId || !membershipTier || !startTime || !endTime) {
      alert("Please fill in all required fields.");
      return;
    }
    try {
      const formattedStartTime = formatDateTime(startTime);
      const formattedEndTime = formatDateTime(endTime);

      const response = await estimateBilling({
        membership_tier: membershipTier,
        start_time: formattedStartTime,
        end_time: formattedEndTime,
      });
      setEstimatedCost(response);
    } catch (error) {
      console.error("Error estimating billing:", error);
      alert(error.message);
    }
  };

  // Generate invoice
  const handleGenerateInvoice = async () => {
    if (!userId || !userEmail || !billingResult?.total) {
      alert("Please fill in all required fields.");
      return;
    }
    try {
      const response = await generateInvoice({
        user_email: userEmail,
        user_id: userId,
        reservation_id: 5, // Static value, adjust accordingly
        total_amount: billingResult.total,
      });
      setInvoiceData(response);
    } catch (error) {
      console.error("Error generating invoice:", error);
      alert(error.message);
    }
  };

  return (
    <div className="container billing-management">
      <h1>Billing Management</h1>

      <div className="billing-form">
        <label>User ID:</label>
        <input
          value={userId}
          onChange={(e) => setUserId(e.target.value)}
          placeholder="Enter user ID"
        />

        <label>Membership Tier:</label>
        <select
          value={membershipTier}
          onChange={(e) => setMembershipTier(e.target.value)}
        >
          <option value="Basic">Basic</option>
          <option value="Premium">Premium</option>
          <option value="VIP">VIP</option>
        </select>

        <label>Start Time (YYYY-MM-DD HH:MM:SS):</label>
        <input
          type="datetime-local"
          value={startTime}
          onChange={(e) => setStartTime(e.target.value)}
        />

        <label>End Time (YYYY-MM-DD HH:MM:SS):</label>
        <input
          type="datetime-local"
          value={endTime}
          onChange={(e) => setEndTime(e.target.value)}
        />

        <label>User Email (For Invoice):</label>
        <input
          type="email"
          value={userEmail}
          onChange={(e) => setUserEmail(e.target.value)}
          placeholder="Enter user email"
        />
      </div>

      <div className="billing-buttons">
        <button onClick={handleCalculateBilling}>Calculate Billing</button>
        <button onClick={handleEstimateBilling}>Estimate Billing</button>
        <button onClick={handleGenerateInvoice}>Generate Invoice</button>
      </div>

      {billingResult && (
        <div className="billing-result">
          <h2>Billing Result</h2>
          <p>Cost: ${billingResult.cost}</p>
          <p>Discount: ${billingResult.discount}</p>
          <p>Total: ${billingResult.total}</p>
        </div>
      )}

      {estimatedCost && (
        <div className="billing-result">
          <h2>Estimated Cost</h2>
          <p>Cost: ${estimatedCost.cost}</p>
          <p>Discount: ${estimatedCost.discount}</p>
          <p>Total: ${estimatedCost.total}</p>
        </div>
      )}

      {invoiceData && (
        <div className="billing-result">
          <h2>Invoice</h2>
          <p>{invoiceData.message}</p>
        </div>
      )}
    </div>
  );
}

export default BillingManagement;
