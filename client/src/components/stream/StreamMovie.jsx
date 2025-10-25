import ReactPlayer from "react-player";
import { useParams } from "react-router-dom";
import "./StreamMovie.css";

export default function StreamMovie() {
  let params = useParams();
  let yt_id = params.yt_id;

  return (
    <div className="react-player-container">
      {yt_id != null ? (
        <ReactPlayer
          controls="true"
          url={`http://www.youtube.com/watch?v=${yt_id}`}
          width="100%"
          height="100%"
        />
      ) : (
        <h2>Oops!!</h2>
      )}
    </div>
  );
}
