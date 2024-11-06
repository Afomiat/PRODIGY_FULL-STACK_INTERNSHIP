import { all } from 'redux-saga/effects';
import { watchAuth } from './sagas/authSaga';

export default function* rootSaga() {
    yield all([
        watchAuth(),
        // Add other sagas here
    ]);
}
