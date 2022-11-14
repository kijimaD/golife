import React from "react";

const pako = require("pako");

const compress = (input: string) => {
  const deflated = pako.deflate(input, { to: "string" }); // 出力はバイナリ文字列の配列 Uint8Array(18) [120, 156, 171, 86, 74, 76, 76, 84, 178, 50, 52, 174, 5, 0, 17, 7, 2, 254]
  const base64Encoded = btoa(deflated); // 64進数へ
  return base64Encoded;
};

// atob(base64文字列) は通常の文字列を返す

export function decompress(input: string) {
  try {
    const base64Decoded = atob(input); // バイナリのコンマ区切りの文字列 "120,156,171,86,74,76,76,84,178,50,52,174,5,0,17,7,2,254"
    // inflateは引数にuint8Arrayを取るので、バイナリ文字列を変換する必要がある。
    const arr = base64Decoded.split(",");
    const uintArray = new Uint8Array(arr.map(Number));
    return pako.inflate(uintArray, { to: "string" });
  } catch {
    return;
  }
}

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
          "http://kijimad.github.io/golife?s=" +
          compress(JSON.stringify(squares)) +
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
