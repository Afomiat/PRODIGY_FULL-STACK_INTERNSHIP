import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Login from './components/Auth/Login';
import Signup from './components/Auth/Signup';
import OTPVerification from './components/Auth/OTPVerification';
import ProtectedRoute from './components/ProtectedRoute';

const App = () => {
    return (
        <Router>
            <Switch>
                <Route path="/login" component={Login} />
                <Route path="/signup" component={Signup} />
                <Route path="/otp-verification" component={OTPVerification} />
                {/* Protected routes can be defined here */}
                {/* <ProtectedRoute path="/dashboard" component={Dashboard} allowedRoles={['admin']} /> */}
            </Switch>
        </Router>
    );
};

export default App;
