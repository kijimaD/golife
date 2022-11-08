import React, { useState, useEffect } from "react";

const ApiFetch = () => {
  var form = new FormData();
  form.append("Debug", "true");
  form.append("GenCap", "3");
  form.append("InitialWorld", "●●○\n○○○\n○○○");

  const [worlds, setWorld] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8888/world/create", {
      method: "POST",
      body: form,
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data.Worlds);
        setWorld(data.Worlds);
      });
  }, []);

  worlds.map((world) => console.log(world["Cells"]));

  return (
    <div>
      {worlds.map((world: any) => (
        <ul>
          {world["Cells"].map((cell: any) => (
            <li>{cell["IsLive"]}</li>
          ))}
        </ul>
      ))}
    </div>
  );
};

export default ApiFetch;
