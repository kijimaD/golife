import React from "react";
import Square from "./Square";

export const Board = () => {
  function renderSquare(bool: boolean) {
    return <Square IsLive={bool} />;
  }

  return (
    <>
      <div className="board-row">
        {renderSquare(true)}
        {renderSquare(false)}
        {renderSquare(true)}
      </div>
      <div className="board-row">
        {renderSquare(false)}
        {renderSquare(false)}
        {renderSquare(false)}
      </div>
      <div className="board-row">
        {renderSquare(false)}
        {renderSquare(true)}
        {renderSquare(true)}
      </div>
    </>
  );
};

export default Board;
