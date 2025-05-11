import { createSlice } from '@reduxjs/toolkit';
import { Server } from '../types';

interface ServerState {
  servers: Server[];
  loading: boolean;
  error: string | null;
}

const initialState: ServerState = {
  servers: [],
  loading: false,
  error: null,
};

const serversSlice = createSlice({
  name: 'servers',
  initialState,
  reducers: {},
});

export default serversSlice.reducer; 