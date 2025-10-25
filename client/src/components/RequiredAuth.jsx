import useAuth from "../hooks/useAuth";
import { Navigate, Outlet, useLocation } from "react-router-dom";

export default function RequiredAuth() {
  const { auth } = useAuth();
  const location = useLocation();

  return auth ? <Outlet /> : <Navigate to="/login" state={{ form: location }} replace />;
}
