import { all } from 'redux-saga/effects';
import { serversSaga } from './servers';
import { metricsSaga } from './metrics';
import { alertsSaga } from './alerts';
import { optimizationSaga } from './optimization';

export function* rootSaga() {
    yield all([
        serversSaga(),
        metricsSaga(),
        alertsSaga(),
        optimizationSaga(),
    ]);
} 