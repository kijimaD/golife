import React from "react";
import Square from "./Square";

function Board({ squares, setSquares }: { squares: any; setSquares: any }) {
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
      squares.map((square: boolean, index: number) =>
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
        {renderSquare(3)}
        {renderSquare(4)}
      </div>
      <div className="board-row">
        {renderSquare(5)}
        {renderSquare(6)}
        {renderSquare(7)}
        {renderSquare(8)}
        {renderSquare(9)}
      </div>
      <div className="board-row">
        {renderSquare(10)}
        {renderSquare(11)}
        {renderSquare(12)}
        {renderSquare(13)}
        {renderSquare(14)}
      </div>
      <div className="board-row">
        {renderSquare(15)}
        {renderSquare(16)}
        {renderSquare(17)}
        {renderSquare(18)}
        {renderSquare(19)}
      </div>
      <div className="board-row">
        {renderSquare(20)}
        {renderSquare(21)}
        {renderSquare(21)}
        {renderSquare(23)}
        {renderSquare(24)}
      </div>
    </>
  );
}

export default Board;
