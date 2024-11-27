import React, { useState, useEffect } from 'react';
import { getAllAttendanceRecords } from '../../api/authApi';

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
    <div>
      <h2>Attendance Records</h2>
      <ul>
        {attendanceRecords.map((record, index) => (
          <li key={index}>
            {record.user_id} - Clock In: {formatDateTime(record.clock_in)} - Clock Out: {record.clock_out && record.clock_out !== "0001-01-01T00:00:00Z" ? formatDateTime(record.clock_out) : 'N/A'}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AttendanceTracking;
