import axios from 'axios';

const API_URL = 'http://localhost:5000/api/auth'; // Adjust the URL based on your backend

export const login = async (credentials) => {
    const response = await axios.post(`${API_URL}/login`, credentials);
    return response.data;
};

export const signup = async (data) => {
    const response = await axios.post(`${API_URL}/signup`, data);
    return response.data;
};

export const verifyOTP = async (data) => {
    const response = await axios.post(`${API_URL}/verify`, data);
    return response.data;
};

export const refreshToken = async () => {
    const response = await axios.post(`${API_URL}/refresh`);
    return response.data;
};
