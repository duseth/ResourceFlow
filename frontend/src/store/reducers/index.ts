import { combineReducers } from '@reduxjs/toolkit';
import { serversReducer } from './servers';
import { metricsReducer } from './metrics';
import { alertsReducer } from './alerts';
import { optimizationReducer } from './optimization';
import { uiReducer } from './ui';

export const rootReducer = combineReducers({
    servers: serversReducer,
    metrics: metricsReducer,
    alerts: alertsReducer,
    optimization: optimizationReducer,
    ui: uiReducer,
}); 