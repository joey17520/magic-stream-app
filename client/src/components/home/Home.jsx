import { useEffect, useState } from "react";
import axiosClient from "../../api/axiosConfig";
import Movies from "../movies/Movies";

export default function Home({ updateMovieReview }) {
  const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(false);
  const [message, SetMessage] = useState();

  useEffect(() => {
    const fetchMovies = async () => {
      setLoading(true);
      SetMessage("");
      try {
        const response = await axiosClient.get("/movies");
        setMovies(response.data);

        if (response.data.length === 0) {
          SetMessage("There are currently no movies available");
        }
      } catch (error) {
        console.error("Error Fetching movies: ", error);
        SetMessage("Error Fetching movies");
      } finally {
        setLoading(false);
      }
    };

    fetchMovies();
  }, []);

  if (loading) {
    return <h2>Loading...</h2>;
  }

  return <Movies movies={movies} message={message} updateMovieReview={updateMovieReview} />;
}
