import React from "react";
import { History, World, Cell } from "../types/History";

const LIVECHAR = "●";
const DEADCHAR = "○";

export const Slide = ({ history }: { history?: History }) => {
  return (
    <div>
      {history &&
        history.Worlds.map((world: World, i: number) => (
          <ul>
            <li>{i}世代</li>
            {world.Cells.map((cell: Cell, j: number) => (
              <span className="Stage">
                {cell.IsLive ? LIVECHAR : DEADCHAR}
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

export default Slide;
