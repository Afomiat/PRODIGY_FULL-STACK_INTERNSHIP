import React, { useState, useEffect } from 'react';
import EmployeeProfile from './EmployeeProfile';
import AttendanceTracking from './AttendanceTracking';
import { createEmployee, updateEmployee, deleteEmployee, getAllEmployees, getAllAttendanceRecords } from '../../api/authApi';
import './AdminDashboard.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEnvelope, faUser, faLock, faEye, faEyeSlash, faHandPaper } from '@fortawesome/free-solid-svg-icons';
const AdminDashboard = () => {
  const [employees, setEmployees] = useState([]);
  const [editingEmployee, setEditingEmployee] = useState(null);
  const [attendanceRecords, setAttendanceRecords] = useState([]);
  const [activeSection, setActiveSection] = useState('employees'); 
  const [isSidebarCollapsed, setIsSidebarCollapsed] = useState(false); 

  useEffect(() => {
    fetchEmployees();
    fetchAttendanceRecords();
  }, []);

  const fetchEmployees = async () => {
    try {
      const allEmployees = await getAllEmployees();
      setEmployees(allEmployees);
    } catch (error) {
      console.error('Error fetching employees:', error);
    }
  };

  const fetchAttendanceRecords = async () => {
    try {
      const records = await getAllAttendanceRecords();
      setAttendanceRecords(records);
    } catch (error) {
      console.error('Error fetching attendance records:', error);
    }
  };

  const handleSaveEmployee = async (employee) => {
    const formattedEmployee = {
      username: employee.username,
      password: employee.password,
      email: employee.email,
      role: employee.role || 'Employee'
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
    setActiveSection('addEmployee');
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

  const renderSection = () => {
    switch (activeSection) {
      case 'attendance':
        return <AttendanceTracking />;
      case 'addEmployee':
        return <EmployeeProfile onSave={handleSaveEmployee} initialData={editingEmployee} />;
      default:
        return (
          <div className="emp-list">
            <h2 className='emp-title'>Employee List</h2>
            <ul>
              {employees
                .filter(employee => employee.role === 'EMPLOYEE')
                .map((employee, index) => (
                  <li key={index}>

                    <div className="names">
                      <FontAwesomeIcon icon={faUser} className="icon-user" />
                      {employee.name} {employee.username} {employee.jobTitle}

                    </div>
                    <button className='edit-but' onClick={() => handleEditEmployee(employee)}>Edit</button>
                    <button className='delete-but' onClick={() => handleDeleteEmployee(employee._id)}>Delete</button>
                  </li>
                ))}
            </ul>
          </div>
        );
    }
  };

  return (
    <div className="admin-dashboard">
      <div className={`sidebar ${isSidebarCollapsed ? 'collapsed' : ''}`}>
        <button className="collapse-button" onClick={() => setIsSidebarCollapsed(!isSidebarCollapsed)}>
          {isSidebarCollapsed ? '➤' : '➤'}
        </button>
        {!isSidebarCollapsed && (
          <div className="sidebar-content">
            <button onClick={() => setActiveSection('employees')}>Employee List</button>
            <button onClick={() => setActiveSection('attendance')}>Attendance</button>
            <button onClick={() => setActiveSection('addEmployee')}>Add Employee</button>
          </div>
        )}
      </div>
      <div className="main-content">
        {renderSection()}
      </div>
    </div>
  );
};

export default AdminDashboard;
