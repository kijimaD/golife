import React from "react";

// パラメータを圧縮する
const Signature = ({
  squares,
  width,
}: {
  squares: boolean[];
  width: number;
}) => {
  return (
    <>
      <a
        href={
          "https://kijimad.github.io/golife?s=" +
          JSON.stringify(squares) +
          "&w=" +
          width
        }
      >
        この世界へのリンク
      </a>
    </>
  );
};

export default Signature;
