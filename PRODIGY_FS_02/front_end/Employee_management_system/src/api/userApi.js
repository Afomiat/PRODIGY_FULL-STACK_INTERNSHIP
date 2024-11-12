import axios from 'axios';

const API_URL = 'http://localhost:5000/api/users'; // Adjust the URL based on your backend

export const fetchAllUsers = async () => {
    const response = await axios.get(`${API_URL}/all_users`);
    return response.data;
};
