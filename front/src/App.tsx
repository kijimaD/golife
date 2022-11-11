import React, { useState, useEffect } from "react";
import "./App.css";
import Loading from "./components/Loading";
import AppHeader from "./layouts/AppHeader";
import { History } from "./types/History";
import Slide from "./components/Slide";
import Anim from "./components/Anim";
import Board from "./components/Board";
import Signature, { decompress } from "./components/Signature";

function App() {
  const [isLoading, setIsLoading] = useState(false);
  const genRef = React.createRef<HTMLInputElement>();
  const [history, setHistory] = useState<History>();
  const [width, setWidth] = useState(5);
  const [squares, setSquares] = useState(Array(width ** 2).fill(false));
  const DEFAULT_GENCAP = 50;

  const squareParam = getParam("s");
  const widthParam = getParam("w");
  useEffect(() => {
    if (squareParam && widthParam) {
      setSquares(JSON.parse(decompress(squareParam)));
      setWidth(Number(widthParam));
    }
  }, [squareParam, widthParam]);

  // TODO: 環境変数で本番用、開発用を切り替えたい
  // let fetchUrl = "http://localhost:8888/world/create";
  let fetchUrl = "https://kd-golife.herokuapp.com/world/create";

  // バックエンドの入力が文字列なので、変換が必要。json入力を受け付けられるようになったら消せる
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

  const incrementWidth = () => setWidth((prevWidth) => prevWidth + 1);
  const decrementWidth = () => setWidth((prevWidth) => prevWidth - 1);

  return (
    <div className="App">
      <AppHeader />
      <button
        onClick={() => {
          incrementWidth();
          setSquares(Array((width + 1) ** 2).fill(false));
          /* スコープに入った時点でwidth確定し前のままなので、+1する必要がある  */
        }}
        className="App-submit"
        type="button"
      >
        ➕
      </button>
      <button
        onClick={() => {
          decrementWidth();
          setSquares(Array((width - 1) ** 2).fill(false));
        }}
        className="App-submit"
        type="button"
      >
        ➖
      </button>
      <Signature squares={squares} width={width} />
      <label className="App-lb">初期世界 ■=生 □=死</label>
      <Board squares={squares} setSquares={setSquares} width={width} />
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

// パラメータを取得する
// (stackoverflowからコピった) react routerを導入してなくて自前で取るしかないため
function getParam(name: string, url?: string) {
  if (!url) url = window.location.href;
  name = name.replace(/[[\]]/g, "\\$&");
  var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
    results = regex.exec(url);
  if (!results) return null;
  if (!results[2]) return "";
  return decodeURIComponent(results[2].replace(/\+/g, " "));
}

export default App;
