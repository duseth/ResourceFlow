import { configureStore } from '@reduxjs/toolkit';
import serversReducer from './slices/serversSlice';
import alertsReducer from './slices/alertsSlice';
import optimizationsReducer from './slices/optimizationsSlice';

export const store = configureStore({
  reducer: {
    servers: serversReducer,
    alerts: alertsReducer,
    optimizations: optimizationsReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch; 