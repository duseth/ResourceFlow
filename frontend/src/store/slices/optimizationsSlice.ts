import { createSlice } from '@reduxjs/toolkit';
import { Optimization } from '../types';

interface OptimizationState {
  optimizations: Optimization[];
  loading: boolean;
  error: string | null;
}

const initialState: OptimizationState = {
  optimizations: [],
  loading: false,
  error: null,
};

const optimizationsSlice = createSlice({
  name: 'optimizations',
  initialState,
  reducers: {},
});

export default optimizationsSlice.reducer; 