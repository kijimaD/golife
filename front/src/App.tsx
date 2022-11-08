import React, { useState } from "react";
import logo from "./rms.png";
import tomato from "./tomato.png";
import "./App.css";

function App() {
  const worldRef = React.createRef<HTMLTextAreaElement>();
  const genRef = React.createRef<HTMLInputElement>();

  type Cell = {
    IsLive: boolean;
  };

  type World = {
    Cells: Cell[];
  };

  type Config = {
    Debug: boolean;
    GenCap: number;
    Row: number;
    Col: number;
  };

  type History = {
    Worlds: World[];
    Configs: Config;
  };

  const [history, setHistory] = useState<History>();
  // TODO: 環境変数で本番用、開発用を切り替えたい
  // let fetch_url = "http://localhost:8888/world/create";
  let fetch_url = "https://kd-golife.herokuapp.com/world/create";

  function handleSubmit(e: any) {
    e.preventDefault();

    var form = new FormData();
    form.append("Debug", "true");
    if (worldRef.current) {
      form.append("InitialWorld", worldRef.current.value);
    }
    if (genRef.current) {
      form.append("GenCap", genRef.current.value);
    }

    fetch(fetch_url, {
      method: "POST",
      body: form,
    })
      .then((res) => res.json())
      .then((data) => {
        setHistory(data);
      });
  }

  const LIVECHAR = "●";
  const DEADCHAR = "○";

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
          初期状態に基づいて、各世代の世界を返します。APIサーバ直アクセスの例↓
          <pre>
            curl -X POST -d $'Debug=trueGenCap=10&InitialWorld=●●○\n○○○\n○○○'
            http://kd-golife.herokuapp.com/world/create
          </pre>
        </p>
      </header>

      <li>世界の縦横の幅が同じでないとバグる</li>

      <form>
        <label className="App-lb">初期世界 ●=生きている ○=死んでいる</label>
        <textarea
          ref={worldRef}
          className="App-textarea"
          defaultValue="○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○●○●○○○○○○○○&#13;○○○○○○○○○○●●●○○○○○○○&#13;○○○○○○○○○●○●○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○●○○○○○○○○&#13;○○○○○○○○○○○○●○○○○○○○&#13;○○○○○○○○○○○○○●○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○"
        />
        <label className="App-lb">生成数</label>
        <input ref={genRef} type="number" defaultValue="100" />
        <label className="App-lb"></label>
        <form>
          <button onClick={handleSubmit}>🚀創造</button>
        </form>
        {history &&
          history.Worlds.map((world: World, i: number) => (
            <ul>
              <li>{i}世代</li>
              {world["Cells"].map((cell: Cell, j: number) => (
                <span className="Stage">
                  {cell["IsLive"] ? LIVECHAR : DEADCHAR}
                  {(j % history.Configs.Row) - history.Configs.Row + 1 === 0 ? (
                    <br />
                  ) : (
                    ""
                  )}
                </span>
              ))}
            </ul>
          ))}
      </form>
    </div>
  );
}

export default App;
