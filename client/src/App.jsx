import "./App.css";
import Home from "./components/home/Home";
import Header from "./components/header/Header";
import { Route, Routes, useNavigate } from "react-router-dom";
import Register from "./components/register/Register";
import Login from "./components/login/Login";
import Layout from "./components/Layout";
import RequiredAuth from "./components/RequiredAuth";
import Recommended from "./components/recommended/Recommended";
import Review from "./components/review/Review";
import axiosConfig from "./api/axiosConfig";
import useAuth from "./hooks/useAuth";
import StreamMovie from "./components/stream/StreamMovie";

function App() {
  const navigate = useNavigate();
  const { auth, setAuth } = useAuth();

  const updateMovieReview = (imdb_id) => {
    navigate(`/review/${imdb_id}`);
  };

  // 注销功能
  const handleLogout = async () => {
    try {
      const response = await axiosConfig.post("/logout", { userid: auth.user_id });
      alert(response.data?.message);
      setAuth(null);
      localStorage.removeItem("user");
      console.log("user logged out");
    } catch (error) {
      console.log("Error logging out: ", error);
    }
  };

  return (
    <>
      <Header handleLogout={handleLogout} />
      <Routes path="/" element={<Layout />}>
        <Route path="/" element={<Home updateMovieReview={updateMovieReview} />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route element={<RequiredAuth />}>
          <Route path="/recommended" element={<Recommended />} />
          <Route path="/review/:imdb_id" element={<Review />} />
          <Route path="/stream/:yt_id" element={<StreamMovie />} />
        </Route>
      </Routes>
    </>
  );
}

export default App;
