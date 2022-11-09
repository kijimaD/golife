import React, { useState } from "react";
import "./App.css";
import Loading from "./components/Loading";
import AppHeader from "./layouts/AppHeader";
import { History, Cell, World } from "./utils/History";

function App() {
  const [isLoading, setIsLoading] = useState(false);
  const worldRef = React.createRef<HTMLTextAreaElement>();
  const genRef = React.createRef<HTMLInputElement>();

  const [history, setHistory] = useState<History>();
  // TODO: 環境変数で本番用、開発用を切り替えたい
  // let fetch_url = "http://localhost:8888/world/create";
  let fetch_url = "https://kd-golife.herokuapp.com/world/create";

  function handleSubmit(e: any) {
    e.preventDefault();
    setIsLoading(true);

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
      .then((res) => {
        setIsLoading(false);
        return res.json();
      })
      .then((data) => {
        setIsLoading(false);
        setHistory(data);
      });
  }

  const LIVECHAR = "●";
  const DEADCHAR = "○";

  return (
    <div className="App">
      <AppHeader />
      <form>
        <label className="App-lb"></label>
        <button onClick={handleSubmit} className="App-submit">
          🚀創造
        </button>
        <label className="App-lb">生成数</label>
        <input ref={genRef} type="number" defaultValue="100" />
        <label className="App-lb">初期世界 ●=生きている ○=死んでいる</label>
        <textarea
          ref={worldRef}
          className="App-textarea"
          defaultValue="○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○●○●○○○○○○○○&#13;○○○○○○○○○○●●●○○○○○○○&#13;○○○○○○○○○●○●○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○●○○○○○○○○&#13;○○○○○○○○○○○○●○○○○○○○&#13;○○○○○○○○○○○○○●○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○&#13;○○○○○○○○○○○○○○○○○○○○"
        />
        {isLoading && <Loading />}
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
