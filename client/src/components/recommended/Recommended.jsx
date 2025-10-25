import { useEffect, useState } from "react";
import useAxiosPrivate from "../../hooks/useAxiosPrivate";
import Movies from "../../components/movies/Movies";

export default function Recommended() {
  const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState();
  const axiosPrivate = useAxiosPrivate();

  useEffect(() => {
    const fetchRecommendedMovies = async () => {
      setLoading(true);
      setMessage("");
      try {
        const response = await axiosPrivate.get("/recommendedmovies");
        console.log(response.data);
        setMovies(response.data);
      } catch (error) {
        console.error("Error fetching recommended movies: ", error);
      } finally {
        setLoading(false);
      }
    };

    fetchRecommendedMovies();
  }, []);

  if (loading) {
    return <h2>Loading...</h2>;
  }

  return (
    <>
      <Movies movies={movies} message={message} />
    </>
  );
}
