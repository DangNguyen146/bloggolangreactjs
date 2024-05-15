import React, { useState, useEffect } from 'react';

function Countdown({ redirectUrl }) {
  const [countdown, setCountdown] = useState(3);

  useEffect(() => {
    const countdownTimer = setInterval(() => {
      setCountdown(prevCountdown => prevCountdown - 1);
    }, 1000);

    if (countdown === 0) {
      clearInterval(countdownTimer);
      window.location.href = redirectUrl;
    }

    return () => clearInterval(countdownTimer);
  }, [countdown, redirectUrl]);

  return (
    <div className="countdown">
      <h1>Redirecting in {countdown} seconds...</h1>
    </div>
  );
}

export default Countdown;