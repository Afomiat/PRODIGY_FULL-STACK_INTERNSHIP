import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { loginAsync } from '../../redux/slices/authSlice';
import { useNavigate, Link } from 'react-router-dom';

function Login() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { error, user } = useSelector((state) => state.auth);

  const handleLogin = () => {
    dispatch(loginAsync({ email, password }));
  };

  // Redirect based on role once logged in
  useEffect(() => {
    if (user) {
      const role = user.role;
      console.log("Navigating user based on role:", role); // Debugging
      if (role === 'ADMIN') {
        navigate('/admin/dashboard');
      } else if (role === 'EMPLOYEE') {
        console.log("in to employee dashboard"); // Debugging
        navigate('/employee/dashboard');
      }
    }
  }, [user, navigate]);

  return (
    <div>
      <h2>Login</h2>
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
      />
      <button onClick={handleLogin}>Login</button>
      {error && <p>{error.message}</p>}
      <p>Don't have an account? <Link to="/signup">Sign up</Link></p>
    </div>
  );
}

export default Login;
