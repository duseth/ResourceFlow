import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { OptimizationRecommendation, OptimizationState } from '../../types';

const initialState: OptimizationState = {
    recommendations: [],
    loading: false,
    error: null,
};

const optimizationSlice = createSlice({
    name: 'optimization',
    initialState,
    reducers: {
        fetchRecommendationsRequest: (state, action: PayloadAction<string>) => {
            state.loading = true;
            state.error = null;
        },
        fetchRecommendationsSuccess: (state, action: PayloadAction<OptimizationRecommendation[]>) => {
            state.recommendations = action.payload;
            state.loading = false;
        },
        fetchRecommendationsFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
        applyRecommendationRequest: (state, action: PayloadAction<string>) => {
            state.loading = true;
            state.error = null;
        },
        applyRecommendationSuccess: (state, action: PayloadAction<OptimizationRecommendation>) => {
            const index = state.recommendations.findIndex((r) => r.id === action.payload.id);
            if (index !== -1) {
                state.recommendations[index] = action.payload;
            }
            state.loading = false;
        },
        applyRecommendationFailure: (state, action: PayloadAction<string>) => {
            state.loading = false;
            state.error = action.payload;
        },
    },
});

export const {
    fetchRecommendationsRequest,
    fetchRecommendationsSuccess,
    fetchRecommendationsFailure,
    applyRecommendationRequest,
    applyRecommendationSuccess,
    applyRecommendationFailure,
} = optimizationSlice.actions;

export const optimizationReducer = optimizationSlice.reducer; 