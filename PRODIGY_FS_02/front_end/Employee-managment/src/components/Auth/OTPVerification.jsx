import React, { useState, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useLocation, useNavigate } from 'react-router-dom';
import { verifyOTPAsync } from '../../redux/slices/authSlice';
import { checkIn } from '../../api/authApi';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEnvelope, faCheck } from '@fortawesome/free-solid-svg-icons';
import './Otp.css';

function OTPVerification() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const location = useLocation();
  const [otp, setOtp] = useState(new Array(6).fill(''));
  const inputRefs = useRef([]);

  const { status, error } = useSelector((state) => state.auth);
  const email = location.state?.email;

  const handleChange = (element, index) => {
    if (isNaN(element.value)) return;

    const newOtp = [...otp];
    newOtp[index] = element.value;
    setOtp(newOtp);

    if (element.value && element.nextSibling) {
      element.nextSibling.focus();
    }
  };

  const handleKeyDown = (e, index) => {
    if (e.key === 'Backspace' && !otp[index] && index > 0) {
      inputRefs.current[index - 1].focus();
    }
  };

  const handleSubmit = () => {
    const otpCode = otp.join('');
    if (!email) return;
    dispatch(verifyOTPAsync({ email, otp: otpCode })).then((action) => {
      if (verifyOTPAsync.fulfilled.match(action)) {
        localStorage.setItem('access_token', action.payload.access_token); // Ensure token is stored here
        handleClockIn();
        navigate('/employee/dashboard');
      }
    });
  };

  const handleClockIn = async () => {
    try {
      const token = localStorage.getItem('access_token');
      console.log('Using access token for clock-in:', token);
      await checkIn();
      console.log('Clocked in successfully after OTP verification');
    } catch (error) {
      console.error('Error clocking in after OTP verification:', error);
    }
  };

  return (
    <div className="otp-container">
      <div className="icon">
        <FontAwesomeIcon icon={faEnvelope} className="icon-maill" />
      </div>
      <div className="title">
        <span className="title-otp">OTP</span> Verification
      </div>
      <p className="first-p">
        Please check your email
      </p>
      <p>We have sent a code to your email address</p>
      <div className="otp-inputs">
        {otp.map((data, index) => (
          <input
            key={index}
            type="text"
            maxLength="1"
            value={data}
            onChange={(e) => handleChange(e.target, index)}
            onKeyDown={(e) => handleKeyDown(e, index)}
            ref={(el) => (inputRefs.current[index] = el)}
            className="otp-input"
            aria-label={`OTP input ${index + 1}`}
          />
        ))}
      </div>
      <button className="otp-button" onClick={handleSubmit}>
        Verify 
      </button>
      {status === 'failed' && (
        <p className="error-message">{error?.error || 'Verification failed'}</p>
      )}
    </div>
  );
}

export default OTPVerification;
