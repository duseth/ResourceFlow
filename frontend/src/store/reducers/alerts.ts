import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Alert, AlertRule, AlertsState } from '../../types';

const initialState: AlertsState = {
    items: [],
    rules: [],
    loading: false,
    error: null,
};

const alertsSlice = createSlice({
    name: 'alerts',
    initialState,
    reducers: {
        fetchAlertsRequest: (state) => {
            state.loading = true;
            state.error = null;
        },
        fetchAlertsSuccess: (state, action: PayloadAction<Alert[]>) => {
            state.items = action.payload;
            state.loading = false;
        },
        fetchAlertsFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        fetchAlertRulesRequest: (state) => {
            state.loading = true;
            state.error = null;
        },
        fetchAlertRulesSuccess: (state, action: PayloadAction<AlertRule[]>) => {
            state.rules = action.payload;
            state.loading = false;
        },
        fetchAlertRulesFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        createAlertRuleRequest: (state, action: PayloadAction<Omit<AlertRule, 'id'>>) => {
            state.loading = true;
            state.error = null;
        },
        createAlertRuleSuccess: (state, action: PayloadAction<AlertRule>) => {
            state.rules.push(action.payload);
            state.loading = false;
        },
        createAlertRuleFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        updateAlertRuleRequest: (state, action: PayloadAction<AlertRule>) => {
            state.loading = true;
            state.error = null;
        },
        updateAlertRuleSuccess: (state, action: PayloadAction<AlertRule>) => {
            const index = state.rules.findIndex((r) => r.id === action.payload.id);
            if (index !== -1) {
                state.rules[index] = action.payload;
            }
            state.loading = false;
        },
        updateAlertRuleFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        deleteAlertRuleRequest: (state, action: PayloadAction<string>) => {
            state.loading = true;
            state.error = null;
        },
        deleteAlertRuleSuccess: (state, action: PayloadAction<string>) => {
            state.rules = state.rules.filter((r) => r.id !== action.payload);
            state.loading = false;
        },
        deleteAlertRuleFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
    },
});

export const {
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
} = alertsSlice.actions;

export const alertsReducer = alertsSlice.reducer; 