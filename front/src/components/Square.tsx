import React from "react";

// isLive: boolean
export const Square = ({
  isLive,
  click,
}: {
  isLive: boolean;
  click: () => void;
}) => {
  return (
    <>
      {isLive ? (
        <button className="square live" onClick={click}></button>
      ) : (
        <button className="square death" onClick={click}></button>
      )}
    </>
  );
};

export default Square;
