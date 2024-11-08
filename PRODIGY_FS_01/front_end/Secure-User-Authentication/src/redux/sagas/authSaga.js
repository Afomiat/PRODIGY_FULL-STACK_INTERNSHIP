import { call, put, takeLatest } from 'redux-saga/effects';
import { setUser, setToken } from '../slices/authSlice';
import { login, signup, verifyOTP } from '../../api/authApi';

function* handleLogin(action) {
  try {
    const { user, token } = yield call(login, action.payload);
    yield put(setUser(user));
    yield put(setToken(token));
  } catch (error) {
    console.error('Login failed:', error);
  }
}

function* handleSignup(action) {
  try {
    const response = yield call(signup, action.payload);
    console.log(response.message);
    if (action.callback) action.callback();
  } catch (error) {
    console.error('Signup failed:', error);
  }
}

function* handleVerifyOTP(action) {
  try {
    const response = yield call(verifyOTP, action.payload);
    console.log(response.message);
    if (action.callback) action.callback();
  } catch (error) {
    console.error('OTP Verification failed:', error);
  }
}

export default function* authSaga() {
  yield takeLatest('auth/login', handleLogin);
  yield takeLatest('auth/signup', handleSignup);
  yield takeLatest('auth/verifyOTP', handleVerifyOTP);
}
