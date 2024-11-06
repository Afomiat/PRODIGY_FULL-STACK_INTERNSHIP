import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { login, signup, verifyOTP, refreshToken } from '../../api/authApi';

export const loginUser = createAsyncThunk('auth/login', async (credentials) => {
    const response = await login(credentials);
    return response;
});

export const signupUser = createAsyncThunk('auth/signup', async (userData) => {
    const response = await signup(userData);
    return response;
});

export const verifyOtp = createAsyncThunk('auth/verifyOtp', async (data) => {
    const response = await verifyOTP(data);
    return response;
});

const authSlice = createSlice({
    name: 'auth',
    initialState: {
        isAuthenticated: false,
        role: null,
        accessToken: null,
    },
    reducers: {
        setUserRole(state, action) {
            state.role = action.payload;
        },
        logout(state) {
            state.isAuthenticated = false;
            state.role = null;
            state.accessToken = null;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(loginUser.fulfilled, (state, action) => {
                state.isAuthenticated = true;
                state.accessToken = action.payload.accessToken; // Assume this is returned
                state.role = action.payload.role; // Assume this is returned
            })
            .addCase(signupUser.fulfilled, (state, action) => {
                // Handle signup success
            })
            .addCase(verifyOtp.fulfilled, (state, action) => {
                // Handle OTP verification success
            });
    },
});

export const { setUserRole, logout } = authSlice.actions;
export default authSlice.reducer;
