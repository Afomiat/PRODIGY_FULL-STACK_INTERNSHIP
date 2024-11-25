import React, { useState, useEffect } from 'react';
import EmployeeProfile from './EmployeeProfile';
import AttendanceTracking from './AttendanceTracking';
import { createEmployee, updateEmployee, deleteEmployee, getAllEmployees } from '../../api/authApi';
import './AdminDashboard.css';

const AdminDashboard = () => {
  const [employees, setEmployees] = useState([]);
  const [editingEmployee, setEditingEmployee] = useState(null);

  const fetchEmployees = async () => {
    try {
      const allEmployees = await getAllEmployees();
      setEmployees(allEmployees);
    } catch (error) {
      console.error('Error fetching employees:', error);
    }
  };
  
  useEffect(() => {
    fetchEmployees();
  }, []);
  
  const handleSaveEmployee = async (employee) => {
    const formattedEmployee = {
      username: employee.username, 
      password: employee.password,
      email: employee.email,
      role: employee.role || "Employee" 
    };
  
    console.log('Saving employee:', formattedEmployee);
  
    try {
      if (editingEmployee) {
        await updateEmployee(editingEmployee._id, formattedEmployee);
      } else {
        await createEmployee(formattedEmployee);
      }
      fetchEmployees();
      setEditingEmployee(null);
    } catch (error) {
      console.error('Error saving employee:', error);
    }
  };
  
  

  const handleEditEmployee = (employee) => {
    setEditingEmployee(employee);
  };
  

  const handleDeleteEmployee = async (id) => {
    try {
      console.log('Deleting employee with ID:', id); 
      await deleteEmployee(id);
      fetchEmployees();
    } catch (error) {
      console.error('Error deleting employee:', error);
    }
  };
  
  

  return (
    <div>
      <p className='p-ad'> <span className='span-ad'>Welcome</span> to the Admin Dashboard</p>
      <EmployeeProfile onSave={handleSaveEmployee} initialData={editingEmployee} />
      <AttendanceTracking />
      <h2>Employee List</h2>
      <ul>
        {employees.map((employee, index) => (
          <li key={index}>
            {employee.name} - {employee.username} - {employee.jobTitle}
            <button onClick={() => handleEditEmployee(employee)}>Edit</button>
            <button onClick={() => handleDeleteEmployee(employee._id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AdminDashboard;
