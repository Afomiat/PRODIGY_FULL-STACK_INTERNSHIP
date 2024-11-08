import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { verifyOTPAsync } from '../../redux/slices/authSlice';

function OTPVerification() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [otp, setOtp] = useState('');
  const { status, error } = useSelector((state) => state.auth);

  const handleSubmit = () => {
    console.log("Sending data:", { email, otp });  // Log request data
    dispatch(verifyOTPAsync({ email, otp })).then((action) => {
      if (verifyOTPAsync.fulfilled.match(action)) {
        navigate('/success');  // Redirect to success page after successful OTP verification
      }
    });
  };

  return (
    <div>
      <h2>OTP Verification</h2>
      <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="Enter Email" />
      <input type="text" value={otp} onChange={(e) => setOtp(e.target.value)} placeholder="Enter OTP" />
      <button onClick={handleSubmit}>Verify OTP</button>
      {status === 'failed' && <p>{error?.error || 'Verification failed'}</p>}
    </div>
  );
}

export default OTPVerification;
