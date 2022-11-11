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
      <hr />
      <a href="https://ja.wikipedia.org/wiki/%E3%83%A9%E3%82%A4%E3%83%95%E3%82%B2%E3%83%BC%E3%83%A0">
        ライフゲームのルール
      </a>
      <p>プリセット</p>
      <ul>
        ■固体物体
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywxNjEsMzksODksODIsODQsMTAsMzcsMTM2LDgwLDEzNSw4NywxMjgsMTI0LDE3OSw3LDEsMjUsMTEsMCw5Miw5MSwxNDAsMTM5&w=8">
            蜂の巣
          </a>
        </li>
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywyNSwyMDQsMTAwLDczLDgxLDQxLDE0OCwzMiw5MCw1Myw2LDE0NywxMDQsNjEsMTE2LDM2LDk5LDEsMjAxLDM5LDE0MCwyMTQ=&w=8">
            ボート
          </a>
        </li>
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywxNjEsNTMsODksODIsODQsMTAsMzcsMjQwLDE3MCwxOTIsOTYsMTQ2LDk5LDE4LDg1LDIwMSw4OCwwLDIwMiw1MiwxMDcsODI=&w=7">
            船
          </a>
        </li>
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywxNjEsMzksODksODIsODQsMTAsMzcsMTM2LDgwLDEzNSw4NywxMjgsMTI0LDIxMywxNjQsMTg1LDEzMiw2LDEwMCw0NCwwLDg2LDg2LDEzOSwyNDU=&w=8">
            池
          </a>
        </li>
        ■振動子
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywxNjEsMTMyLDQ0LDQxLDQyLDY5LDM4LDQwLDU0LDQ4LDIyLDAsNCw1Nyw1NSwxMQ==&w=5">
            ブリンカー
          </a>
        </li>
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywyNSwxMzgsMTAwLDczLDgxLDQxLDUwLDY1LDE2MSwxNzgsOTcsNzQsMTk4LDIsMCwyMTAsMTYsMTc4LDUw&w=9">
            ヒキガエル
          </a>
        </li>
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywxNjEsNjMsODksODIsODQsMTAsMzcsMTY4LDE2MywxNTQsMjE4LDIzMCwyMDksMTUyLDE0MCw1LDAsODYsODYsMTM5LDI0NQ==&w=8">
            ビーコン
          </a>
        </li>
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywxNjEsMzksODksODIsODQsNzQsMTY2LDM4LDE4OCw1OCw5LDE3MSwxNjAsMTcwLDEzMSw0MCwzOCw5OSwxLDgxLDk5LDE0MCwxMzk=&w=8">
            時計
          </a>
        </li>
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywzMywxMzQsNDQsNDEsNDIsMTMzLDE4LDE5NiwyMzUsMjUsNzIsMTE1LDEwNSwyMjksNzQsMTE0LDIyMCwxMzksMTc0LDcsMTc1LDExOCwxNjIsMjA1LDE5OCwxMDYsNDIsMTQsMjIxLDIwLDQsNDgsMzAsMjcsMTM2LDE4MCwxNTYsNjgsMjMxLDE0NCwxMDksNDIsNTMsMTYyLDEwNiwzMiwxNDcsMjE0LDIwOCw0OCwxMTksMTMyLDE0NSwxNzcsMCwxNDYsMTUwLDIzMiwxNzg=&w=15">
            時計Ⅱ
          </a>
        </li>
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywyNSwxNzQsMTAwLDczLDgxLDQxLDc4LDEsMTIsNTcsMTYyLDEzLDE5NiwxMTAsOCw2NiwxMzgsNDQsMTQ3LDczLDExNiwyMCwxMzcsMTE4LDE2LDIzNSw0NiwxMCwxNTYsNzgsMjQ3LDE4NCwzNiwyMjMsMTUxLDI0NCw4LDExMywyMzYsMjM4LDM0LDE0OCwxMTYsNDAsOCwyMCwxMjQsMTA1LDEyNSwxNDgsMjgsMTkzLDEwMCw0NCwwLDQ2LDIxNiw0NywxMTk=&w=16">
            銀河
          </a>
        </li>
        ■移動物体
        <li>
          <a href="https://kijimad.github.io/golife/?s=MTIwLDE1NiwxMzksNzgsNzUsMjA0LDQxLDc4LDIxMywxNjEsNTMsODksODIsODQsNzQsNTMsMjI5LDk2LDczLDQsNjUsOTksMTY3LDE5OSwyLDAsMzUsMTQwLDEwNywxNTc=&w=7">
            グライダー
          </a>
        </li>
      </ul>
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
