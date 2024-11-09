import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useLocation, useNavigate } from 'react-router-dom';
import { verifyOTPAsync } from '../../redux/slices/authSlice';

function OTPVerification() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const location = useLocation();
  const [otp, setOtp] = useState('');
  const { status, error } = useSelector((state) => state.auth);

  const email = location.state?.email; // Retrieve email from Signup page

  const handleSubmit = () => {
    if (!email) return; // Ensure email is available
    dispatch(verifyOTPAsync({ email, otp })).then((action) => {
      if (verifyOTPAsync.fulfilled.match(action)) {
        navigate('/employee/dashboard');  // Redirect to success page after successful OTP verification
      }
    });
  };

  return (
    <div>
      <h2>OTP Verification</h2>
      <input type="text" value={otp} onChange={(e) => setOtp(e.target.value)} placeholder="Enter OTP" />
      <button onClick={handleSubmit}>Verify OTP</button>
      {status === 'failed' && <p>{error?.error || 'Verification failed'}</p>}
    </div>
  );
}

export default OTPVerification;
