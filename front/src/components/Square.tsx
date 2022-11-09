import React, { useState } from "react";

export const Square = ({ IsLive }: { IsLive: boolean }) => {
  const [live, setLive] = useState(IsLive);

  return (
    <>
      {live ? (
        <button className="square live" onClick={() => setLive(false)}></button>
      ) : (
        <button className="square death" onClick={() => setLive(true)}></button>
      )}
    </>
  );
};

export default Square;
