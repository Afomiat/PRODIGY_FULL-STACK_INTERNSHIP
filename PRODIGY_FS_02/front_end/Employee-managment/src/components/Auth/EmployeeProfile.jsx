import React, { useState, useEffect } from 'react';

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
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        placeholder="Username"
        required
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
        required
      />
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
        required
      />
      <input
        type="text"
        value={role}
        onChange={(e) => setRole(e.target.value)}
        placeholder="Role"
      />
      <button type="submit">Save</button>
    </form>
  );
};

export default EmployeeProfile;
