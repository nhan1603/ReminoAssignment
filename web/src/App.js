import React, { useState, useEffect } from 'react';
import Login from './components/Login';
import VideoShare from './components/VideoShare';
import SharedVideoList from './components/SharedVideoList';

function App() {
  const [user, setUser] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const token = localStorage.getItem('authToken');
    if (token) {
      setUser({ token });
    }
    setIsLoading(false);
  }, []);

  const handleLogin = (userData) => {
    localStorage.setItem('authToken', userData.token);
    setUser(userData);
  };

  const handleLogout = () => {
    localStorage.removeItem('authToken');
    setUser(null);
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="App">
      <h1>Video Sharing App</h1>
      {user ? (
        <>
          <button onClick={handleLogout}>Logout</button>
          <VideoShare user={user} />
        </>
      ) : (
        <Login onLogin={handleLogin} />
      )}
      <SharedVideoList />
    </div>
  );
}

export default App;