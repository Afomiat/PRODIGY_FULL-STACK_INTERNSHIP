import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { verifyOTP as verifyOTPApi, signup as signupApi } from '../../api/authApi';

export const verifyOTP = createAsyncThunk('auth/verifyOTP', async (otp, { rejectWithValue }) => {
  try {
    const response = await verifyOTPApi(otp);
    return response;
  } catch (error) {
    return rejectWithValue(error.response.data);
  }
});

export const signup = createAsyncThunk('auth/signup', async (userInfo, { rejectWithValue }) => {
  try {
    const response = await signupApi(userInfo);
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
      .addCase(verifyOTP.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(verifyOTP.fulfilled, (state) => {
        state.status = 'succeeded';
      })
      .addCase(verifyOTP.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload;
      })
      .addCase(signup.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(signup.fulfilled, (state) => {
        state.status = 'succeeded';
      })
      .addCase(signup.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload;
      });
  },
});

export const { setUser, setToken } = authSlice.actions;
export default authSlice.reducer;
