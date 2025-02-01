import { call, put, takeLatest } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { api } from '../../services/api';
import {
    fetchLiveMetricsRequest,
    fetchLiveMetricsSuccess,
    fetchLiveMetricsFailure,
    fetchHistoricalMetricsRequest,
    fetchHistoricalMetricsSuccess,
    fetchHistoricalMetricsFailure,
} from '../reducers/metrics';

function* fetchLiveMetrics(action: PayloadAction<string>) {
    try {
        const response = yield call(api.getLiveMetrics, action.payload);
        yield put(fetchLiveMetricsSuccess(response.data));
    } catch (error) {
        yield put(fetchLiveMetricsFailure(error.message));
    }
}

function* fetchHistoricalMetrics(action: PayloadAction<{ serverId: string; startDate: string; endDate: string }>) {
    try {
        const response = yield call(
            api.getHistoricalMetrics,
            action.payload.serverId,
            action.payload.startDate,
            action.payload.endDate
        );
        yield put(fetchHistoricalMetricsSuccess(response.data));
    } catch (error) {
        yield put(fetchHistoricalMetricsFailure(error.message));
    }
}

export function* metricsSaga() {
    yield takeLatest(fetchLiveMetricsRequest.type, fetchLiveMetrics);
    yield takeLatest(fetchHistoricalMetricsRequest.type, fetchHistoricalMetrics);
} 