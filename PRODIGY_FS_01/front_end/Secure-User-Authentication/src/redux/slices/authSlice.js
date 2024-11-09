import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { login, signup, verifyOTP } from '../../api/authApi';

export const verifyOTPAsync = createAsyncThunk('auth/verifyOTP', async (otpData, { rejectWithValue }) => {
  try {
    const response = await verifyOTP(otpData);
    return response;
  } catch (error) {
    console.error('Verification error:', error);
    return rejectWithValue(error.response.data);
  }
});

export const signupAsync = createAsyncThunk('auth/signup', async (userInfo, { rejectWithValue }) => {
  try {
    const response = await signup(userInfo);
    return response;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

export const loginAsync = createAsyncThunk('auth/login', async (credentials, { rejectWithValue }) => {
  try {
    const response = await login(credentials);
    return response;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

const authSlice = createSlice({
  name: 'auth',
  initialState: { user: null, token: null, status: null, error: null },
  reducers: {
    setUser: (state, action) => {
      state.user = action.payload;
    },
    setToken: (state, action) => {
      state.token = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(verifyOTPAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(verifyOTPAsync.fulfilled, (state) => {
        state.status = 'succeeded';
      })
      .addCase(verifyOTPAsync.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload;
      })
      .addCase(signupAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(signupAsync.fulfilled, (state) => {
        state.status = 'succeeded';
      })
      .addCase(signupAsync.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload;
      })
      .addCase(loginAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(loginAsync.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.user = action.payload; // Ensure user state is updated correctly
        state.token = action.payload.accessToken;
      })
      .addCase(loginAsync.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload;
      });
  },
});

export const { setUser, setToken } = authSlice.actions;

export default authSlice.reducer;
