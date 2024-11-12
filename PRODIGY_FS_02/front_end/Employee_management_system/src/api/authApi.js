import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Your actual backend URL

// Create a function to refresh the access token using the refresh token
export const refreshToken = async () => {
  try {
    const response = await axios.post(`${API_BASE_URL}/refresh`, null, {
      withCredentials: true, // Ensures that cookies (including the refresh token) are sent with the request
    });

    return response.data.access_token; // Return the new access token
  } catch (error) {
    console.error('Failed to refresh token:', error);
    throw new Error('Could not refresh token');
  }
};

// Example function to make a request that requires a valid access token
export const getUserData = async () => {
  try {
    const accessToken = localStorage.getItem('access_token');
    const response = await axios.get(`${API_BASE_URL}/user`, {
      headers: {
        Authorization: `Bearer ${accessToken}`, // Send the access token in the Authorization header
      },
      withCredentials: true, // Ensures cookies are sent for this request
    });

    return response.data; // Return user data if successful
  } catch (error) {
    if (error.response && error.response.status === 401) {
      // Token expired or invalid, refresh the token
      const newAccessToken = await refreshToken(); // Get new access token

      // Save the new access token
      localStorage.setItem('access_token', newAccessToken);

      // Retry the original request with the new access token
      return getUserData();
    }

    // If it's not a 401 error, throw the error
    throw error;
  }
};

// Function to handle login
export const login = async (credentials) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/login`, credentials, {
      withCredentials: true, // Ensure cookies are sent with the request
    });

    // Save the access token (and refresh token, if needed) in local storage or handle them as required
    localStorage.setItem('access_token', response.data.access_token);

    return response.data;
  } catch (error) {
    console.error('Login failed:', error);
    throw error;
  }
};

// Function to handle signup
export const signup = async (userInfo) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/signup`, userInfo, {
      withCredentials: true, // Ensure cookies are sent with the request
    });

    localStorage.setItem('access_token', response.data.access_token);
    return response.data;
  } catch (error) {
    console.error('Signup failed:', error);
    throw error;
  }
};

// Function to verify OTP
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
