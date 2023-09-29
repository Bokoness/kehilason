import { Navigate, Outlet } from "react-router-dom";

function ProtectedRout({ authorized }) {
  return authorized ? <Outlet /> : <Navigate to="/" />;
}

export default ProtectedRout;
