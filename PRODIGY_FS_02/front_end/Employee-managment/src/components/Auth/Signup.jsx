import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate, Link } from 'react-router-dom';
import { signupAsync } from '../../redux/slices/authSlice';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEnvelope, faUser, faLock, faEye, faEyeSlash, faHandPaper } from '@fortawesome/free-solid-svg-icons';
import './Signup.css';
import signImage from '../../assets/signup.avif';

function Signup() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false); 
  const { status, error } = useSelector((state) => state.auth);

  const handleSubmit = () => {
    dispatch(signupAsync({ email, username, password })).then((action) => {
      if (signupAsync.fulfilled.match(action)) {
        localStorage.setItem('access_token', action.payload.access_token);
        navigate('/verify', { state: { email } });
      }
    });
  };

  return (
    <>
      <div className="welcom"><FontAwesomeIcon icon={faHandPaper} size="2x" className='hand-wave' /> Welcome!</div>
      <div className='signup-container'>
        <img src={signImage} alt='signup' className='signup-image' />
        <div className="line"></div>
        <div className='signup-box'>
          <h2 className='signup-title'>Signup</h2>
          <div className="signup-inner">
            <div className="signup-input-container">
              <FontAwesomeIcon icon={faEnvelope} className="signup-icon-mail" />
              <input
                className='signup-input'
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="Email"
                id="signup-email"
              />
            </div>
            <div className="signup-input-container">
              <FontAwesomeIcon icon={faUser} className="signup-icon-user" />
              <input
                className='signup-input'
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Username"
                id="signup-username"
              />
            </div>
            <div className="signup-input-container">
              <FontAwesomeIcon icon={faLock} className="signup-icon-pass" />
              <input
                className='signup-input'
                type={showPassword ? "text" : "password"}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Password"
                id="signup-password"
              />
              <FontAwesomeIcon
                icon={showPassword ? faEyeSlash : faEye}
                className="signup-toggle-password"
                onClick={() => setShowPassword(!showPassword)}
              />
            </div>
            <button className='signup-button' onClick={handleSubmit}>Signup</button>
            {status === 'failed' && <p className="signup-error">{error?.error || 'Signup failed'}</p>}
            <p className='signup-p'>Already have an account? <Link to="/login" className='signup-link'>Login</Link></p>
          </div>
        </div>
      </div>
    </>
  );
}

export default Signup;
