import React, { useState } from 'react';

function VideoShare({ onShareVideo }) {
  const [videoUrl, setVideoUrl] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    // TODO: Extract video ID from URL and implement sharing logic
    onShareVideo(videoUrl);
    setVideoUrl('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={videoUrl}
        onChange={(e) => setVideoUrl(e.target.value)}
        placeholder="YouTube Video URL"
        required
      />
      <button type="submit">Share Video</button>
    </form>
  );
}

export default VideoShare;
