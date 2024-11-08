import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { verifyOTP } from '../../redux/slices/authSlice';

function OTPVerification() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [otp, setOtp] = useState('');
  const { status, error } = useSelector((state) => state.auth);

  const handleSubmit = () => {
    dispatch(verifyOTP({ email, otp })).then((action) => {
      if (verifyOTP.fulfilled.match(action)) {
        navigate('/success');
      }
    });
  };

  return (
    <div>
      <h2>OTP Verification</h2>
      <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="Enter Email" />
      <input type="text" value={otp} onChange={(e) => setOtp(e.target.value)} placeholder ="Enter OTP" />
      <button onClick={handleSubmit}>Verify OTP</button>
      {status === 'failed' && <p>{error?.error || 'Verification failed'}</p>}
    </div>
  );
}

export default OTPVerification;
