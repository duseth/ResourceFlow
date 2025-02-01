import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Metric, MetricsState, ResourceUsage } from '../../types';

const initialState: MetricsState = {
    current: {},
    historical: {},
    loading: false,
    error: null,
};

const metricsSlice = createSlice({
    name: 'metrics',
    initialState,
    reducers: {
        fetchLiveMetricsRequest: (state, action: PayloadAction<string>) => {
            state.loading = true;
            state.error = null;
        },
        fetchLiveMetricsSuccess: (
            state,
            action: PayloadAction<{ serverId: string; metrics: ResourceUsage }>
        ) => {
            state.current[action.payload.serverId] = action.payload.metrics;
            state.loading = false;
        },
        fetchLiveMetricsFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        fetchHistoricalMetricsRequest: (
            state,
            action: PayloadAction<{ serverId: string; from: string; to: string }>
        ) => {
            state.loading = true;
            state.error = null;
        },
        fetchHistoricalMetricsSuccess: (
            state,
            action: PayloadAction<{ serverId: string; metrics: Metric[] }>
        ) => {
            state.historical[action.payload.serverId] = action.payload.metrics;
            state.loading = false;
        },
        fetchHistoricalMetricsFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
    },
});

export const {
    fetchLiveMetricsRequest,
    fetchLiveMetricsSuccess,
    fetchLiveMetricsFailure,
    fetchHistoricalMetricsRequest,
    fetchHistoricalMetricsSuccess,
    fetchHistoricalMetricsFailure,
} = metricsSlice.actions;

export const metricsReducer = metricsSlice.reducer; 