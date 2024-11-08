// // src/utils/tokenHelpers.js
// export const setAccessToken = (token) => localStorage.setItem('accessToken', token);
// export const getAccessToken = () => localStorage.getItem('accessToken');
// export const removeAccessToken = () => localStorage.removeItem('accessToken');

// export const isTokenExpired = (token) => {
//     if (!token) return true;

//     const payload = JSON.parse(atob(token.split('.')[1]));
//     return payload.exp * 1000 < Date.now();
// };
