import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { loginAsync } from '../../redux/slices/authSlice';
import { useNavigate, Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEnvelope, faLock, faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons';
import './Login.css';

function Login() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const { error, user } = useSelector((state) => state.auth);

  const handleLogin = () => {
    dispatch(loginAsync({ email, password }));
  };

  useEffect(() => {
    if (user) {
      const role = user.role;
      if (role === 'ADMIN') {
        navigate('/admin/dashboard');
      } else if (role === 'EMPLOYEE') {
        navigate('/employee/dashboard');
      }
    }
  }, [user, navigate]);

  return (
    <div className='login-container'>
      <div className='login-box'>
        <h2 className='login-title'>Login</h2>
        <div className="login-inner">
          <div className="input-container">
            <FontAwesomeIcon icon={faEnvelope} className="icon-mail" />
            <input
              className='login-input'
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="Email"
              id="email"
            />
          </div>

          <div className="input-container">
            <FontAwesomeIcon icon={faLock} className="icon-pass" />
            <input
              className='login-input'
              type={showPassword ? "text" : "password"}
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="Password"
              id="password"
            />
            <FontAwesomeIcon
              icon={showPassword ? faEyeSlash : faEye}
              className="toggle-password"
              onClick={() => setShowPassword(!showPassword)}
            />
          </div>

          <button className='login-button' onClick={handleLogin}>Login</button>
          {error && <p className="login-error">{error.message}</p>}
          <p>Don't have an account? <Link to="/signup" className='login-link'>Sign up</Link></p>
        </div>
      </div>
    </div>
  );
}

export default Login;
