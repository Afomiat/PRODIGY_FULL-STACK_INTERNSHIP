import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { signup } from '../../redux/slices/authSlice';

function Signup() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const { status, error } = useSelector((state) => state.auth);

  const handleSubmit = () => {
    dispatch(signup({ email, username, password })).then((action) => {
      if (signup.fulfilled.match(action)) {
        navigate('/verify');
      }
    });
  };

  return (
    <div>
      <h2>Signup</h2>
      <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="Email" />
      <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} placeholder="Username" />
      <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" />
      <button onClick={handleSubmit}>Signup</button>
      {status === 'failed' && <p>{error}</p>}
    </div>
  );
}

export default Signup;
