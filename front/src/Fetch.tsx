import React, { useState, useEffect } from "react";

const ApiFetch = () => {
  var form = new FormData();
  // ここをフォームから取ってくる
  form.append("Debug", "true");
  form.append("GenCap", "10");
  form.append("InitialWorld", "●○○○○\n○○○●○\n○○●○○\n○○●○○\n○○●○○");

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

  useEffect(() => {
    fetch("http://localhost:8888/world/create", {
      method: "POST",
      body: form,
    })
      .then((res) => res.json())
      .then((data) => {
        setHistory(data);
      });
  }, []);

  const LIVECHAR = "●";
  const DEADCHAR = "○";

  return (
    <div>
      {history &&
        history.Worlds.map((world: World, i: number) => (
          <ul>
            <li>{i}</li>
            {world["Cells"].map((cell: Cell, j: number) => (
              <span>
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
    </div>
  );
};

export default ApiFetch;
