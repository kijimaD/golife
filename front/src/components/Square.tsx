import React from "react";

export const Square = ({ IsLive }: { IsLive: boolean }) => {
  return (
    <>
      {IsLive ? (
        <button className="square live"></button>
      ) : (
        <button className="square death"></button>
      )}
    </>
  );
};

export default Square;
