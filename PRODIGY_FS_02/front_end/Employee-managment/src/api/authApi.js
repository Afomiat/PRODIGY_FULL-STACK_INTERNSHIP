import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/'; 

const api = axios.create({
  baseURL: API_BASE_URL,
  withCredentials: true,
});

api.interceptors.request.use(
  async (config) => {
    let accessToken = localStorage.getItem('access_token');
    if (!accessToken) {
      accessToken = await refreshToken();
      localStorage.setItem('access_token', accessToken);
    }

    config.headers.Authorization = `Bearer ${accessToken}`;
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);


export const refreshToken = async () => {
  try {
    const response = await api.post('/refresh');
    return response.data.access_token; 
  } catch (error) {
    console.error('Failed to refresh token:', error);
    throw new Error('Could not refresh token');
  }
};

export const getUserData = async () => {
  try {
    const response = await api.get('/user');
    return response.data; 
  } catch (error) {
    if (error.response && error.response.status === 401) {
      const newAccessToken = await refreshToken(); 
      localStorage.setItem('access_token', newAccessToken);
      return getUserData();
    }
    throw error;
  }
};

export const login = async (credentials) => {
  try {
    const response = await api.post('/login', credentials);
    localStorage.setItem('access_token', response.data.access_token);
    return response.data;
  } catch (error) {
    console.error('Login failed:', error);
    throw error;
  }
};

export const signup = async (userInfo) => {
  try {
    const response = await api.post('/signup', userInfo);
    localStorage.setItem('access_token', response.data.access_token);
    return response.data;
  } catch (error) {
    console.error('Signup failed:', error);
    throw error;
  }
};

export const verifyOTP = async ({ email, otp }) => {
  try {
    const response = await api.post('/verify', { email, otp });
    return response.data;
  } catch (error) {
    console.error('OTP verification failed:', error);
    throw error;
  }
};

export const createEmployee = async (employeeData) => {
  const response = await api.post('/employees/create_employee', employeeData);
  return response.data;
};

export const updateEmployee = async (id, employeeData) => {
  const response = await api.put(`/employees/update_employee/${id}`, employeeData);
  return response.data;
};

export const deleteEmployee = async (id) => {
  const response = await api.delete(`/employees/delete_employee/${id}`);
  return response.data;
};

export const getEmployee = async (id) => {
  const response = await api.get(`/employees/get_employee/${id}`);
  return response.data;
};

export const getAllEmployees = async () => {
  const response = await api.get('/employees/get_all_employee');
  console.log(response.data)
  return response.data;
};
