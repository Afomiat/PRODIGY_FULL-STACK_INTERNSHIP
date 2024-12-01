import React, { useState, useEffect } from 'react';
import { getAllAttendanceRecords } from '../../api/authApi';
import './Attendance.css';

const AttendanceTracking = () => {
  const [attendanceRecords, setAttendanceRecords] = useState([]);

  useEffect(() => {
    fetchAttendanceRecords();
  }, []);

  const fetchAttendanceRecords = async () => {
    try {
      const records = await getAllAttendanceRecords();
      console.log("Fetched attendance records: ", records);
      setAttendanceRecords(records);
    } catch (error) {
      console.error('Error fetching attendance records:', error);
    }
  };

  const formatDateTime = (dateTime) => {
    const date = new Date(dateTime);
    return `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
  };

  return (
    <div className='attend'>
      <h2>Attendance Records</h2>
      <div className="title-att">
        <span className='emp'>Employee</span>
        
        <span className='in'>clock-in</span>
        <span className='out'>clock-out</span>
      </div>
      <ul>
        <div className="contener">
          {attendanceRecords.map((record, index) => (
            <ul key={index}>
              <div className="content">
                <div className="attendance-name">
                  {record.username}
                </div>
                <div className="clock">
                  <div className='clock-in'>
                    {formatDateTime(record.clock_in && record.clock_in !== "0001-01-01T00:00:00Z" ? record.clock_in : 'No Record')}
                  </div>
                
                  <div className='clock-out'>
                     {record.clock_out && record.clock_out !== "0001-01-01T00:00:00Z" ? formatDateTime(record.clock_out) : 'No Record'}
                  </div>
                </div>
              </div>
            </ul>
          ))}
        </div>
      </ul>
    </div>
  );
};

export default AttendanceTracking;
