import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { verifyOTP } from '../../redux/slices/authSlice';

const OTPVerification = () => {
    const [otp, setOtp] = useState('');
    const dispatch = useDispatch();

    const handleSubmit = (e) => {
        e.preventDefault();
        dispatch(verifyOTP({ otp }));
    };

    return (
        <form onSubmit={handleSubmit}>
            <input type="text" placeholder="Enter OTP" value={otp} onChange={(e) => setOtp(e.target.value)} required />
            <button type="submit">Verify OTP</button>
        </form>
    );
};

export default OTPVerification;
