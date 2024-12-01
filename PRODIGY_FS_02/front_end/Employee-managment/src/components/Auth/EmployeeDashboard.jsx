import React, { useState } from 'react';
import { checkIn, checkOut } from '../../api/authApi';
import './EmployeeDashboard.css';

const EmployeeDashboard = () => {
  const [status, setStatus] = useState('');
  const [isCheckedIn, setIsCheckedIn] = useState(false);

  const handleCheckIn = async () => {
    try {
      await checkIn();
      setStatus('Clocked in successfully');
      setIsCheckedIn(true);
    } catch (error) {
      console.error('Error clocking in:', error);
      setStatus('Error clocking in');
    }
  };

  const handleCheckOut = async () => {
    try {
      await checkOut();
      setStatus('Clocked out successfully');
      setIsCheckedIn(false);
    } catch (error) {
      console.error('Error clocking out:', error);
      setStatus('Error clocking out');
    }
  };

  return (
    <div>
      <p className='p-emp'><span className='span-emp'>Welcome</span> to the Employee Dashboard</p>
      <div className="employee-but">
        <button className='but-in' onClick={handleCheckIn} disabled={isCheckedIn}>Clock In</button>
        <button className='but-out'onClick={handleCheckOut} disabled={!isCheckedIn}>Clock Out</button>
        <p>{status}</p>

      </div>
    </div>
  );
};

export default EmployeeDashboard;
