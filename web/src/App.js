import React, { useState, useEffect } from 'react';
import Login from './components/Login';
import VideoShare from './components/VideoShare';
import SharedVideoList from './components/SharedVideoList';

function App() {
  const [user, setUser] = useState(null);
  const [refreshTrigger, setRefreshTrigger] = useState(0);

  useEffect(() => {
    const token = localStorage.getItem('authToken');
    if (token) {
      setUser({ token });
    }
  }, []);

  const handleLogin = (userData) => {
    localStorage.setItem('authToken', userData.token);
    setUser(userData);
  };

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    setUser(null);
  };

  const refreshSharedVideos = () => {
    setRefreshTrigger(prev => prev + 1);
  };

  return (
    <div className="App">
      <h1>Video Sharing App</h1>
      {user ? (
        <>
          <button onClick={handleLogout}>Logout</button>
          <VideoShare user={user} onVideoShared={refreshSharedVideos} />
        </>
      ) : (
        <Login onLogin={handleLogin} />
      )}
      <SharedVideoList refreshTrigger={refreshTrigger} />
    </div>
  );
}

export default App;