import React from 'react';
import { Route, Navigate } from 'react-router-dom';
import { useSelector } from 'react-redux';

const ProtectedRoute = ({ element: Component, allowedRoles, ...rest }) => {
    const { role, isAuthenticated } = useSelector((state) => state.auth);

    return (
        <Route
            {...rest}
            element={
                !isAuthenticated || !allowedRoles.includes(role) ? (
                    <Navigate to="/unauthorized" />
                ) : (
                    <Component />
                )
            }
        />
    );
};

export default ProtectedRoute;
