import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Update to your backend URL and port

export const login = async (credentials) => {
  const response = await axios.post(`${API_BASE_URL}/login`, credentials, { withCredentials: true });
  return response.data;
};

export const signup = async (userInfo) => {
  const response = await axios.post(`${API_BASE_URL}/signup`, userInfo, { withCredentials: true });
  return response.data;
};

export const verifyOTP = async (otp) => {
  const response = await axios.post(`${API_BASE_URL}/verify`, { otp }, { withCredentials: true });
  return response.data;
};
