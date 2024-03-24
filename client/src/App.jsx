import gopherGif from "./assets/gopher-dance-long-3x.gif";
import oauth2Logo from "./assets/Oauth_logo.svg";
import "./App.css";

function App() {
  const handleLogin = () => {
    window.location.href = "http://localhost:3000/auth/google";
  };

  const handleLogout = () => {
    window.location.href = "http://localhost:3000/logout/google";
  };

  return (
    <>
      <div>
        <a href="https://go.dev" target="_blank">
          <img src={gopherGif} className="logo react" alt="Go logo" />
        </a>
        <a href="https://go.dev" target="_blank">
          <img src={oauth2Logo} className="logo vite" alt="OAuth2 logo" />
        </a>
      </div>
      <h1>Go & OAuth2</h1>
      <div className="card">
        <p>Auth implementation with Go and React</p>
      </div>

      <div>
        <button onClick={handleLogin}>Login with Google</button>
        <button onClick={handleLogout}>Logout</button>
      </div>

      <p className="read-the-docs">OAuth2 with Go</p>
    </>
  );
}

export default App;
