import { call, put, takeLatest } from 'redux-saga/effects';
import { PayloadAction } from '@reduxjs/toolkit';
import { api } from '../../services/api';
import { OptimizationRecommendation } from '../../types';
import {
    fetchRecommendationsRequest,
    fetchRecommendationsSuccess,
    fetchRecommendationsFailure,
    applyRecommendationRequest,
    applyRecommendationSuccess,
    applyRecommendationFailure,
} from '../reducers/optimization';

function* fetchRecommendations(action: PayloadAction<string>) {
    try {
        const response = yield call(api.getOptimizationRecommendations, action.payload);
        yield put(fetchRecommendationsSuccess(response.data));
    } catch (error) {
        yield put(fetchRecommendationsFailure(error.message));
    }
}

function* applyRecommendation(action: PayloadAction<{ serverId: string; recommendationId: string }>) {
    try {
        const response = yield call(
            api.applyOptimizationRecommendation,
            action.payload.serverId,
            action.payload.recommendationId
        );
        yield put(applyRecommendationSuccess(response.data));
    } catch (error) {
        yield put(applyRecommendationFailure(error.message));
    }
}

export function* optimizationSaga() {
    yield takeLatest(fetchRecommendationsRequest.type, fetchRecommendations);
    yield takeLatest(applyRecommendationRequest.type, applyRecommendation);
} 