import React, { useState } from "react";
import "./BillingManagement.css";
import {
  calculateBilling,
  estimateBilling,
  generateInvoice,
} from "../utils/api";

function BillingManagement() {
  const [userId, setUserId] = useState("");
  const [billingResult, setBillingResult] = useState(null);
  const [estimatedCost, setEstimatedCost] = useState(null);
  const [invoiceData, setInvoiceData] = useState(null);

  // Calculate billing
  const handleCalculateBilling = async () => {
    if (!userId) {
      alert("Please enter a User ID.");
      return;
    }
    try {
      const response = await calculateBilling(userId);
      setBillingResult(response);
    } catch (error) {
      console.error("Error calculating billing:", error);
      alert(error.message);
    }
  };

  // Estimate billing
  const handleEstimateBilling = async () => {
    if (!userId) {
      alert("Please enter a User ID.");
      return;
    }
    try {
      const response = await estimateBilling(userId);
      setEstimatedCost(response);
    } catch (error) {
      console.error("Error estimating billing:", error);
      alert(error.message);
    }
  };

  // Generate invoice
  const handleGenerateInvoice = async () => {
    if (!userId) {
      alert("Please enter a User ID.");
      return;
    }
    try {
      const response = await generateInvoice(userId);
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
      </div>

      <div className="billing-buttons">
        <button onClick={handleCalculateBilling}>Calculate Billing</button>
        <button onClick={handleEstimateBilling}>Estimate Billing</button>
        <button onClick={handleGenerateInvoice}>Generate Invoice</button>
      </div>

      {billingResult && (
        <div>
          <h2>Billing Result</h2>
          <pre>{JSON.stringify(billingResult, null, 2)}</pre>
        </div>
      )}
      {estimatedCost && (
        <div>
          <h2>Estimated Cost</h2>
          <pre>{JSON.stringify(estimatedCost, null, 2)}</pre>
        </div>
      )}
      {invoiceData && (
        <div>
          <h2>Invoice Data</h2>
          <pre>{JSON.stringify(invoiceData, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}

export default BillingManagement;
