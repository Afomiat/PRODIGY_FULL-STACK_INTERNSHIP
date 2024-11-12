import { combineReducers } from 'redux';
import authReducer from './slices/authSlice';

const rootReducer = combineReducers({
  auth: authReducer,
  // Add other reducers here
});

export default rootReducer;
