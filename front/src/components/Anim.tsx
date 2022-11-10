import React, { useState, useEffect } from "react";
import { History, Cell } from "../types/History";

const LIVECHAR = "●";
const DEADCHAR = "○";

// 受信時にタイマーを有効化するため、historyは非nil
const Anim = ({ history }: { history: History }) => {
  const [count, setCount] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setCount(count + 1);
    }, 100);
    return () => clearInterval(interval);
  }, [count]);

  return (
    <div>
      <ul>
        <li>
          {count % history.Configs.GenCap}世代{" "}
          {".".repeat(count % history.Configs.GenCap)}
        </li>
        {history.Worlds[count % history.Configs.GenCap].Cells.map(
          (cell: Cell, j: number) => (
            <span className="Stage">
              {cell.IsLive ? LIVECHAR : DEADCHAR}
              {(j % history.Configs.Row) - history.Configs.Row + 1 === 0 ? (
                <br />
              ) : (
                ""
              )}
            </span>
          )
        )}
      </ul>
    </div>
  );
};

export default Anim;
