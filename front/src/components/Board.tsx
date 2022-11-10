import React from "react";
import Square from "./Square";

function Board({
  squares,
  setSquares,
  width,
}: {
  squares: any;
  setSquares: any;
  width: number;
}) {
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

  // 連番二次元配列
  // [[0, 1, 2],
  //  [3, 4, 5],
  //  [6, 7, 8]]
  function twoDimIndex() {
    const renban = [...Array(squares.length)].map((_, i) => i);
    const results = [];
    while (renban.length) results.push(renban.splice(0, width));
    return results;
  }

  return (
    <>
      {twoDimIndex().map((row, ri) => {
        return (
          <div className="board-row">
            {row.map((square, si) => {
              return renderSquare(row[si]);
            })}
          </div>
        );
      })}
    </>
  );
}

export default Board;
