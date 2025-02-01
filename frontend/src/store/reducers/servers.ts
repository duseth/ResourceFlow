import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Server, ServersState } from '../../types';

const initialState: ServersState = {
    items: [],
    selectedServer: null,
    loading: false,
    error: null,
};

const serversSlice = createSlice({
    name: 'servers',
    initialState,
    reducers: {
        fetchServersRequest: (state) => {
            state.loading = true;
            state.error = null;
        },
        fetchServersSuccess: (state, action: PayloadAction<Server[]>) => {
            state.items = action.payload;
            state.loading = false;
        },
        fetchServersFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        selectServer: (state, action: PayloadAction<Server>) => {
            state.selectedServer = action.payload;
        },
        createServerRequest: (state, action: PayloadAction<Omit<Server, 'id'>>) => {
            state.loading = true;
            state.error = null;
        },
        createServerSuccess: (state, action: PayloadAction<Server>) => {
            state.items.push(action.payload);
            state.loading = false;
        },
        createServerFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        updateServerRequest: (state, action: PayloadAction<Server>) => {
            state.loading = true;
            state.error = null;
        },
        updateServerSuccess: (state, action: PayloadAction<Server>) => {
            const index = state.items.findIndex((s) => s.id === action.payload.id);
            if (index !== -1) {
                state.items[index] = action.payload;
            }
            state.loading = false;
        },
        updateServerFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        deleteServerRequest: (state, action: PayloadAction<string>) => {
            state.loading = true;
            state.error = null;
        },
        deleteServerSuccess: (state, action: PayloadAction<string>) => {
            state.items = state.items.filter((s) => s.id !== action.payload);
            state.loading = false;
        },
        deleteServerFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
    },
});

export const {
    fetchServersRequest,
    fetchServersSuccess,
    fetchServersFailure,
    selectServer,
    createServerRequest,
    createServerSuccess,
    createServerFailure,
    updateServerRequest,
    updateServerSuccess,
    updateServerFailure,
    deleteServerRequest,
    deleteServerSuccess,
    deleteServerFailure,
} = serversSlice.actions;

export const serversReducer = serversSlice.reducer; 