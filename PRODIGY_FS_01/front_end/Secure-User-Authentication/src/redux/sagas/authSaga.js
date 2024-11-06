import { takeLatest, call, put } from 'redux-saga/effects';
import { loginUser, signupUser, verifyOtp } from '../slices/authSlice';
import { login, signup, verifyOTP } from '../../api/authApi';

function* handleLogin(action) {
    try {
        const response = yield call(login, action.payload);
        yield put(loginUser.fulfilled(response));
    } catch (error) {
        console.error(error);
        // Handle error
    }
}

function* handleSignup(action) {
    try {
        const response = yield call(signup, action.payload);
        yield put(signupUser.fulfilled(response));
    } catch (error) {
        console.error(error);
        // Handle error
    }
}

function* handleVerifyOtp(action) {
    try {
        const response = yield call(verifyOTP, action.payload);
        yield put(verifyOtp.fulfilled(response));
    } catch (error) {
        console.error(error);
        // Handle error
    }
}

export function* watchAuth() {
    yield takeLatest(loginUser.type, handleLogin);
    yield takeLatest(signupUser.type, handleSignup);
    yield takeLatest(verifyOtp.type, handleVerifyOtp);
}
