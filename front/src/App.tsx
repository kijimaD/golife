import React from "react";
import logo from "./rms.png";
import tomato from "./tomato.png";
import "./App.css";
import ApiFetch from "./Fetch";

function App() {
  const worldRef = React.createRef<HTMLTextAreaElement>();
  const genRef = React.createRef<HTMLInputElement>();

  function handleSubmit(e: any) {
    e.preventDefault();
    console.log("test world", worldRef.current?.value);
    console.log("test gen", genRef.current?.value);
    // refをもとにfetchしたい
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="" alt="logo" />
        <img src={tomato} className="App-logo" alt="logo" />
        <img src={tomato} className="App-logo" alt="logo" />
        <img src={tomato} className="App-logo" alt="logo" />
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
        <textarea
          ref={worldRef}
          className="App-textarea"
          defaultValue="○○●○○&#13;○○○●○&#13;○○●○○&#13;○○●○○&#13;○○●○○"
        />
        <label className="App-lb">生成数</label>
        <input ref={genRef} type="number" defaultValue="100" />
        <label className="App-lb"></label>
        <form>
          <button onClick={handleSubmit}>🚀創造</button>
        </form>
        <ApiFetch />
      </form>
    </div>
  );
}

export default App;
