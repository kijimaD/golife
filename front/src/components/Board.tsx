import React, { useState } from "react";
import Square from "./Square";

function Board() {
  const [squares, setSquares] = useState(Array(9).fill(true));

  function renderSquare(i: number) {
    return (
      <Square
        isLive={squares[i]}
        click={() => {
          flipSquare(i);
        }}
      />
    );
  }

  const flipSquare = (target_i: number) => {
    setSquares(
      squares.map((square, index) =>
        index === target_i ? !squares[index] : squares[index]
      )
    );
  };

  return (
    <>
      <div className="board-row">
        {renderSquare(0)}
        {renderSquare(1)}
        {renderSquare(2)}
      </div>
      <div className="board-row">
        {renderSquare(3)}
        {renderSquare(4)}
        {renderSquare(5)}
      </div>
      <div className="board-row">
        {renderSquare(6)}
        {renderSquare(7)}
        {renderSquare(8)}
      </div>
    </>
  );
}

export default Board;
