// Utility functions for handling tokens

// Save the access token in local storage
export const saveToken = (token) => {
    localStorage.setItem('access_token', token);
  };
  
  // Retrieve the access token from local storage
  export const getToken = () => {
    return localStorage.getItem('access_token');
  };
  
  // Remove the access token from local storage
  export const removeToken = () => {
    localStorage.removeItem('access_token');
  };
  