import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './components/Auth/Login';
import Signup from './components/Auth/Signup';
import OTPVerification from './components/Auth/OTPVerification';
import Success from './components/Auth/Success';
import ProtectedRoute from './components/Auth/ProtectedRoute';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/verify" element={<OTPVerification />} />
        <Route path="/success" element={<Success />} />
        <Route path="/" element={<ProtectedRoute />}>
          <Route path="/" element={<div>Protected Home</div>} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;
