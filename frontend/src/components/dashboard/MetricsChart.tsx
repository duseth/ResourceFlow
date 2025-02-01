import React, { useEffect, useState } from 'react';
import { Box, FormControl, InputLabel, MenuItem, Select } from '@mui/material';
import { Line } from 'react-chartjs-2';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { useSelector } from 'react-redux';
import { RootState, Metric } from '../../types';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

const options = {
    responsive: true,
    plugins: {
        legend: {
            position: 'top' as const,
        },
        title: {
            display: false,
        },
    },
    scales: {
        y: {
            beginAtZero: true,
        },
    },
};

const MetricsChart: React.FC = () => {
    const [selectedMetric, setSelectedMetric] = useState('cpu');
    const metrics = useSelector((state: RootState) => state.metrics.historical);
    const servers = useSelector((state: RootState) => state.servers.items);

    const getChartData = () => {
        const labels = metrics[servers[0]?.id]?.map((m) =>
            new Date(m.timestamp).toLocaleTimeString()
        ) || [];

        return {
            labels,
            datasets: servers.map((server) => ({
                label: server.name,
                data: metrics[server.id]
                    ?.filter((m) => m.type === selectedMetric)
                    .map((m) => m.value) || [],
                borderColor: `hsl(${Math.random() * 360}, 70%, 50%)`,
                tension: 0.4,
            })),
        };
    };

    return (
        <Box>
            <FormControl fullWidth sx={{ mb: 2 }}>
                <InputLabel>Метрика</InputLabel>
                <Select
                    value={selectedMetric}
                    label="Метрика"
                    onChange={(e) => setSelectedMetric(e.target.value)}
                >
                    <MenuItem value="cpu">CPU</MenuItem>
                    <MenuItem value="memory">Memory</MenuItem>
                    <MenuItem value="disk">Disk</MenuItem>
                    <MenuItem value="network">Network</MenuItem>
                </Select>
            </FormControl>
            <Box sx={{ height: 400 }}>
                <Line options={options} data={getChartData()} />
            </Box>
        </Box>
    );
};

export default MetricsChart; 