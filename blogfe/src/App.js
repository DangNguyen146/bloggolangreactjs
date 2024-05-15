import React, { useState, useEffect } from 'react';
import './App.css';
import ProgressBar from 'react-bootstrap/ProgressBar';
import Countdown from './components/Countdown';

function App() {
  const [redirectUrl, setRedirectUrl] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [progress, setProgress] = useState(1);

  useEffect(() => {
    const path = window.location.pathname;

    if (path === '/') {
      setIsLoading(false);
      return; // Kết thúc useEffect nếu không có đường dẫn
    }

    const shortLink = path.substring(1);

    if (shortLink) {
      fetch('https://dangshort.onrender.com/api/telegram', {
        method: 'POST',
        body: JSON.stringify({ path })
      })
      .then(response => response.json())
      .then(data => {
      })
      .catch(error => console.error('Error sending header:', error));
      
      const fetchPromise = fetch(`https://script.google.com/macros/s/AKfycby1kpcTV29jb24xu6QBZe5uPY69JY4QBobsI5LG_w3jwevxqHpt2sIRVLXlArZor-YP/exec?Short=${shortLink}`)
        .then(response => response.json())
        .then(data => {
          if (data.status === true) {
            setRedirectUrl(data.originalLink);
          } else {
            console.log(data.originalLink);
          }
        })
        .catch(error => console.error('Error fetching data:', error));

      const progressInterval = setInterval(() => {
        setProgress(prevProgress => {
          if (prevProgress === 200) {
            clearInterval(progressInterval);
            setIsLoading(false);
            return prevProgress;
          }
          return prevProgress + 1;
        });
      }, 30); // Tăng progress mỗi 30ms

      const countdownPromise = new Promise(resolve => setTimeout(resolve, 3000)); // Đợi 3 giây

      Promise.all([fetchPromise, countdownPromise])
        .then(() => setIsLoading(false)); // Khi cả hai promise đều hoàn thành, setIsLoading(false)
    }
  }, []);

  return (
    <div className="App">
      <div className='container'>
      {isLoading ? (
        <ProgressBar className='ProgressBarChange' now={progress} variant="info" label={`Loading... ${progress}%`} />
      ) : (
        redirectUrl ? <Countdown redirectUrl={redirectUrl} /> : <h1>No valid link found.</h1>
      )}
    </div></div>
  );
}

export default App;
