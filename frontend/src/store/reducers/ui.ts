import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { UIState, Notification } from '../../types';

const initialState: UIState = {
    theme: 'dark',
    sidebarOpen: true,
    notifications: [],
};

const uiSlice = createSlice({
    name: 'ui',
    initialState,
    reducers: {
        toggleTheme: (state) => {
            state.theme = state.theme === 'light' ? 'dark' : 'light';
        },
        toggleSidebar: (state) => {
            state.sidebarOpen = !state.sidebarOpen;
        },
        addNotification: (state, action: PayloadAction<Omit<Notification, 'id'>>) => {
            state.notifications.push({
                ...action.payload,
                id: Date.now().toString(),
            });
        },
        removeNotification: (state, action: PayloadAction<string>) => {
            state.notifications = state.notifications.filter((n) => n.id !== action.payload);
        },
        clearNotifications: (state) => {
            state.notifications = [];
        },
    },
});

export const {
    toggleTheme,
    toggleSidebar,
    addNotification,
    removeNotification,
    clearNotifications,
} = uiSlice.actions;

export const uiReducer = uiSlice.reducer; 