import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; 
export const refreshToken = async () => {
  try {
    const response = await axios.post(`${API_BASE_URL}/refresh`, null, {
      withCredentials: true, 
    });

    return response.data.access_token; 
  } catch (error) {
    console.error('Failed to refresh token:', error);
    throw new Error('Could not refresh token');
  }
};

export const getUserData = async () => {
  try {
    const accessToken = localStorage.getItem('access_token');
    const response = await axios.get(`${API_BASE_URL}/user`, {
      headers: {
        Authorization: `Bearer ${accessToken}`, 
      },
      withCredentials: true, 
    });

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
    const response = await axios.post(`${API_BASE_URL}/login`, credentials, {
      withCredentials: true, 
    });

    localStorage.setItem('access_token', response.data.access_token);

    return response.data;
  } catch (error) {
    console.error('Login failed:', error);
    throw error;
  }
};

export const signup = async (userInfo) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/signup`, userInfo, {
      withCredentials: true, 
    });

    localStorage.setItem('access_token', response.data.access_token);
    return response.data;
  } catch (error) {
    console.error('Signup failed:', error);
    throw error;
  }
};

export const verifyOTP = async ({ email, otp }) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/verify`, { email, otp }, {
      withCredentials: true,
    });

    return response.data;
  } catch (error) {
    console.error('OTP verification failed:', error);
    throw error;
  }
};
