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
  const genRef = React.createRef<HTMLInputElement>();
  const [history, setHistory] = useState<History>();
  const [width, setWidth] = useState(5);
  const [squares, setSquares] = useState(Array(width ** 2).fill(false));
  const DEFAULT_GENCAP = 50;

  // TODO: 環境変数で本番用、開発用を切り替えたい
  // let fetchUrl = "http://localhost:8888/world/create";
  let fetchUrl = "https://kd-golife.herokuapp.com/world/create";

  function squaresToString(squares: any[], width: number) {
    let arr = squares.map((square: boolean, i: number) => {
      if (square === true) {
        return "●";
      } else {
        return "○";
      }
    });

    let arrWithNewline = arr.map((s, i) => {
      if (i % width === width - 1) {
        return s + "\n";
      }
      return s;
    });

    return arrWithNewline.join("");
  }

  function handleSubmit() {
    setIsLoading(true);

    var form = new FormData();
    form.append("Debug", "true");
    form.append("InitialWorld", squaresToString(squares, width));
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
      <label className="App-lb">初期世界 ■=生 □=死</label>
      <Board squares={squares} setSquares={setSquares} />
      <label className="App-lb"></label>
      <label className="App-lb">生成世代数</label>
      <input ref={genRef} type="number" defaultValue={DEFAULT_GENCAP} />
      <button onClick={handleSubmit} className="App-submit" type="button">
        🚀創造
      </button>
      {isLoading && <Loading />}
      {history && <Anim history={history} />}
      <hr />
      <Slide history={history} />
    </div>
  );
}

export default App;
