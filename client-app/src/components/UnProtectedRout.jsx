import React from "react";
import { Navigate, Outlet } from "react-router-dom";

function UnProtectedRout({ authorized }) {
  return !authorized ? <Outlet /> : <Navigate to="/" />;
}

export default UnProtectedRout;
