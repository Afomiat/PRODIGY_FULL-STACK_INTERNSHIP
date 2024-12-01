import React, { useState, useEffect } from 'react';
import './EmployeeProfile.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
 import { faIdBadge } from '@fortawesome/free-solid-svg-icons';

const EmployeeProfile = ({ onSave, initialData }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState(''); // Ensure password field is included
  const [email, setEmail] = useState('');
  const [role, setRole] = useState('Employee');

  useEffect(() => {
    if (initialData) {
      setUsername(initialData.username);
      setPassword(initialData.password);
      setEmail(initialData.email);
      setRole(initialData.role);
    }
  }, [initialData]);

  const handleSubmit = (e) => {
    e.preventDefault();
    onSave({ username, password, email, role });
    setUsername('');
    setPassword('');
    setEmail('');
    setRole('Employee');
  };

  return (
    <form onSubmit={handleSubmit} className='profile-container'>
<FontAwesomeIcon icon={faIdBadge} className="fa-badge" /> {/* Font Awesome badge icon */}      <div className="form-group">
        <label htmlFor="username"></label>
        <input
          type="text"
          id="username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Enter Username"
          required
        />
      </div>

      <div className="form-group">
        <label htmlFor="password"></label>
        <input
          type="password"
          id="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Enter Password"
          required
        />
      </div>

      <div className="form-group">
        <label htmlFor="email"></label>
        <input
          type="email"
          id="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Enter Email"
          required
        />
      </div>

      <div className="form-group">
        <label htmlFor="role"></label>
        <input
          type="text"
          id="role"
          value={role}
          onChange={(e) => setRole(e.target.value)}
          placeholder="Enter Role"
        />
      </div>

      <button type="submit">Add Employee</button>
    </form>
  );
};

export default EmployeeProfile;
