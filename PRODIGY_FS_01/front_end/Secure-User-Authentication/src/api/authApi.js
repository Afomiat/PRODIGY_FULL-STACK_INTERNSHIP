import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Replace with your actual backend URL

export const login = async (credentials) => {
  const response = await axios.post(`${API_BASE_URL}/login`, credentials, { withCredentials: true });
  return response.data;
};

export const signup = async (userInfo) => {
  const response = await axios.post(`${API_BASE_URL}/signup`, userInfo, { withCredentials: true });
  return response.data;
};

export const verifyOTP = async ({ email, otp }) => {
  const response = await axios.post(`${API_BASE_URL}/verify`, { email, otp }, { withCredentials: true });
  return response.data;
};
