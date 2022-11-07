import React from "react";
import logo from "./rms.png";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="" alt="logo" />
        <h2>
          <a
            className="App-link"
            href="https://github.com/kijimaD/golife"
            target="_blank"
            rel="noopener noreferrer"
          >
            ★ライフゲーム★
          </a>
        </h2>
        <p>
          これはトップページです。↓で渡された初期状態に基づいて、各世代の世界を返します。
        </p>
      </header>
      <form>
        <label className="App-lb">初期世界</label>
        <textarea className="App-textarea" defaultValue="●○○&#13;○●○&#13;○●○" />
        <label className="App-lb">世代数</label>
        <input type="number" value="20" />
        <label className="App-lb">🚀</label>
        <input type="submit" />
      </form>
    </div>
  );
}

export default App;
