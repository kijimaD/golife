import React from "react";
import logo from "../rms.png";
import tomato from "../tomato.png";

const AppHeader = () => {
  return (
    <header>
      <img src={logo} className="" alt="logo" />
      <img src={tomato} className="App-logo" alt="logo" />
      <img src={tomato} className="App-logo" alt="logo" />
      <img src={tomato} className="App-logo" alt="logo" />
      <h3>
        <a
          href="https://github.com/kijimaD/golife"
          target="_blank"
          rel="noopener noreferrer"
        >
          ★ライフゲーム★
        </a>
      </h3>
      <p>
        初期状態に基づいて、各世代の世界を返します。APIサーバ直アクセスの例↓
        ⚠Herokuインスタンスがウォームアップしていないと時間がかかります
      </p>
    </header>
  );
};

export default AppHeader;
