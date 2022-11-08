import React, { useState, useEffect } from "react";

const ApiFetch = () => {
  var form = new FormData();
  // ここをフォームから取ってくる
  form.append("Debug", "true");
  form.append("GenCap", "3");
  form.append("InitialWorld", "●○○\n○●○\n○○○");

  const [worlds, setWorlds] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8888/world/create", {
      method: "POST",
      body: form,
    })
      .then((res) => res.json())
      .then((data) => {
        setWorlds(data.Worlds);
      });
  }, []);

  const LIVECHAR = "●";
  const DEADCHAR = "○";

  return (
    <div>
      {worlds.map((world: any, i) => (
        <ul>
          <li>{i}</li>
          {world["Cells"].map((cell: any, j: number) => (
            <span>{cell["IsLive"] ? LIVECHAR : DEADCHAR}</span>
          ))}
        </ul>
      ))}
    </div>
  );
};

export default ApiFetch;
