import React, { useState, useEffect } from 'react';
import Login from './components/Login';
import VideoList from './components/VideoList';
import VideoShare from './components/VideoShare';
import axios from './axiosConfig';

function App() {
  const [user, setUser] = useState(null);
  const [videos, setVideos] = useState([]);

  useEffect(() => {
    // TODO: Fetch videos from API when user is logged in
    if (user) {
      // Simulated API call
      setVideos([
        { id: 1, title: 'Sample Video 1', youtube_video_id: 'dQw4w9WgXcQ' },
        { id: 2, title: 'Sample Video 2', youtube_video_id: 'dQw4w9WgXcQ' },
      ]);
    }
  }, [user]);

  const handleLogin = (username) => {
    setUser({ username });
  };

  const handleShareVideo = async ({ videoUrl, targetEmail }) => {
    try {
      const response = await axios.post('/api/public/v1/share-video', { videoUrl, targetEmail });
      console.log('Video shared successfully:', response.data);
    } catch (error) {
      console.error('Error sharing video:', error);
    }
  };

  if (!user) {
    return <Login onLogin={handleLogin} />;
  }

  return (
    <div>
      <h1>Welcome, {user.username}!</h1>
      <VideoShare onShareVideo={handleShareVideo} />
      <VideoList videos={videos} />
    </div>
  );
}

export default App;
