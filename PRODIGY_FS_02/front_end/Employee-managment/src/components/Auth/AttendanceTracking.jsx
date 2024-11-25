import React, { useState } from 'react';

const AttendanceTracking = () => {
  const [attendance, setAttendance] = useState([]);

  const handleClockIn = () => {
    setAttendance([...attendance, { type: 'Clock-in', time: new Date().toLocaleString() }]);
  };

  const handleClockOut = () => {
    setAttendance([...attendance, { type: 'Clock-out', time: new Date().toLocaleString() }]);
  };

  return (
    <div>
      <button onClick={handleClockIn}>Clock In</button>
      <button onClick={handleClockOut}>Clock Out</button>
      <ul>
        {attendance.map((entry, index) => (
          <li key={index}>{entry.type} at {entry.time}</li>
        ))}
      </ul>
    </div>
  );
};

export default AttendanceTracking;
