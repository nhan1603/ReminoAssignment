import React, { useState } from 'react';
import axios from '../axiosConfig';

function Login({ onLogin }) {
  const [isLogin, setIsLogin] = useState(true);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    try {
      const endpoint = isLogin ? '/api/public/v1/login' : '/api/public/v1/user';
      const response = await axios.post(endpoint, { email, password });
      if (isLogin) {
        onLogin(response.data.user);
      } else {
        setIsLogin(true);
        setError('Account created successfully. Please log in.');
      }
    } catch (err) {
      setError(isLogin ? 'Invalid credentials' : 'Failed to create account');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="email"
        required
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
        required
      />
      {error && <p className="error">{error}</p>}
      <button type="submit">{isLogin ? 'Login' : 'Create Account'}</button>
      <button type="button" onClick={() => setIsLogin(!isLogin)}>
        {isLogin ? 'Need an account?' : 'Already have an account?'}
      </button>
    </form>
  );
}

export default Login;