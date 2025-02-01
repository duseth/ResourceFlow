import { call, put, takeLatest } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { api } from '../../services/api';
import { Alert, AlertRule } from '../../types';
import {
    fetchAlertsRequest,
    fetchAlertsSuccess,
    fetchAlertsFailure,
    fetchAlertRulesRequest,
    fetchAlertRulesSuccess,
    fetchAlertRulesFailure,
    createAlertRuleRequest,
    createAlertRuleSuccess,
    createAlertRuleFailure,
    updateAlertRuleRequest,
    updateAlertRuleSuccess,
    updateAlertRuleFailure,
    deleteAlertRuleRequest,
    deleteAlertRuleSuccess,
    deleteAlertRuleFailure,
} from '../reducers/alerts';

function* fetchAlerts() {
    try {
        const response = yield call(api.getAlerts);
        yield put(fetchAlertsSuccess(response.data));
    } catch (error) {
        yield put(fetchAlertsFailure(error.message));
    }
}

function* fetchAlertRules() {
    try {
        const response = yield call(api.getAlertRules);
        yield put(fetchAlertRulesSuccess(response.data));
    } catch (error) {
        yield put(fetchAlertRulesFailure(error.message));
    }
}

function* createAlertRule(action: PayloadAction<Omit<AlertRule, 'id'>>) {
    try {
        const response = yield call(api.createAlertRule, action.payload);
        yield put(createAlertRuleSuccess(response.data));
    } catch (error) {
        yield put(createAlertRuleFailure(error.message));
    }
}

function* updateAlertRule(action: PayloadAction<AlertRule>) {
    try {
        const response = yield call(api.updateAlertRule, action.payload.id, action.payload);
        yield put(updateAlertRuleSuccess(response.data));
    } catch (error) {
        yield put(updateAlertRuleFailure(error.message));
    }
}

function* deleteAlertRule(action: PayloadAction<string>) {
    try {
        yield call(api.deleteAlertRule, action.payload);
        yield put(deleteAlertRuleSuccess(action.payload));
    } catch (error) {
        yield put(deleteAlertRuleFailure(error.message));
    }
}

export function* alertsSaga() {
    yield takeLatest(fetchAlertsRequest.type, fetchAlerts);
    yield takeLatest(fetchAlertRulesRequest.type, fetchAlertRules);
    yield takeLatest(createAlertRuleRequest.type, createAlertRule);
    yield takeLatest(updateAlertRuleRequest.type, updateAlertRule);
    yield takeLatest(deleteAlertRuleRequest.type, deleteAlertRule);
} 