import React, { useState } from "react";
import "./App.css";
import Loading from "./components/Loading";
import AppHeader from "./layouts/AppHeader";
import { History } from "./types/History";
import Slide from "./components/Slide";
import Anim from "./components/Anim";
import Board from "./components/Board";

function App() {
  const [isLoading, setIsLoading] = useState(false);
  const worldRef = React.createRef<HTMLTextAreaElement>();
  const genRef = React.createRef<HTMLInputElement>();
  const [history, setHistory] = useState<History>();
  const [squares, setSquares] = useState(Array(width ** 2).fill(true));
  const DEFAULT_WORLD =
    "○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○●○●○○○○○○○○\n○○○○○○○○○○●●●○○○○○○○\n○○○○○○○○○●○●○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○●○○○○○○○○\n○○○○○○○○○○○○●○○○○○○○\n○○○○○○○○○○○○○●○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○\n○○○○○○○○○○○○○○○○○○○○";
  const DEFAULT_GENCAP = 50;

  // TODO: 環境変数で本番用、開発用を切り替えたい
  // let fetchUrl = "http://localhost:8888/world/create";
  let fetchUrl = "https://kd-golife.herokuapp.com/world/create";

  function handleSubmit() {
    setIsLoading(true);

    var form = new FormData();
    form.append("Debug", "true");
    if (worldRef.current) {
      form.append("InitialWorld", worldRef.current.value);
    }
    if (genRef.current) {
      form.append("GenCap", genRef.current.value);
    }

    fetch(fetchUrl, {
      method: "POST",
      body: form,
    })
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        setIsLoading(false);
        setHistory(data);
      });
  }

  return (
    <div className="App">
      <AppHeader />
      <Board squares={squares} setSquares={setSquares} />
      <form>
        <label className="App-lb"></label>
        <button onClick={handleSubmit} className="App-submit" type="button">
          🚀創造
        </button>
        <label className="App-lb">生成数</label>
        <input ref={genRef} type="number" defaultValue={DEFAULT_GENCAP} />
        <label className="App-lb">初期世界 ●=生きている ○=死んでいる</label>
        <textarea
          ref={worldRef}
          className="App-textarea"
          defaultValue={DEFAULT_WORLD}
        />
        {isLoading && <Loading />}
        {history && <Anim history={history} />}
        <hr />
        <Slide history={history} />
      </form>
    </div>
  );
}

export default App;
