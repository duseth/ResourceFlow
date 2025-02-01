import { call, put, takeLatest } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { api } from '../../services/api';
import { Server } from '../../types';
import {
    fetchServersRequest,
    fetchServersSuccess,
    fetchServersFailure,
    createServerRequest,
    createServerSuccess,
    createServerFailure,
    updateServerRequest,
    updateServerSuccess,
    updateServerFailure,
    deleteServerRequest,
    deleteServerSuccess,
    deleteServerFailure,
} from '../reducers/servers';

function* fetchServers() {
    try {
        const response = yield call(api.getServers);
        yield put(fetchServersSuccess(response.data));
    } catch (error) {
        yield put(fetchServersFailure(error.message));
    }
}

function* createServer(action: PayloadAction<Omit<Server, 'id'>>) {
    try {
        const response = yield call(api.createServer, action.payload);
        yield put(createServerSuccess(response.data));
    } catch (error) {
        yield put(createServerFailure(error.message));
    }
}

function* updateServer(action: PayloadAction<Server>) {
    try {
        const response = yield call(api.updateServer, action.payload.id, action.payload);
        yield put(updateServerSuccess(response.data));
    } catch (error) {
        yield put(updateServerFailure(error.message));
    }
}

function* deleteServer(action: PayloadAction<string>) {
    try {
        yield call(api.deleteServer, action.payload);
        yield put(deleteServerSuccess(action.payload));
    } catch (error) {
        yield put(deleteServerFailure(error.message));
    }
}

export function* serversSaga() {
    yield takeLatest(fetchServersRequest.type, fetchServers);
    yield takeLatest(createServerRequest.type, createServer);
    yield takeLatest(updateServerRequest.type, updateServer);
    yield takeLatest(deleteServerRequest.type, deleteServer);
} 