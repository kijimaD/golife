import React from "react";
import logo from "../rms.png";
import tomato from "../tomato.png";

const AppHeader = () => {
  return (
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
        ⚠Herokuインスタンスがウォームアップしていないと時間がかかります
      </p>
      <pre>
        curl -X POST -d $'Debug=trueGenCap=10&InitialWorld=●●○\n○○○\n○○○'
        http://kd-golife.herokuapp.com/world/create
      </pre>
    </header>
  );
};

export default AppHeader;
