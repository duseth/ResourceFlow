import React from 'react';
import { Grid, Paper, Typography } from '@mui/material';
import { useSelector } from 'react-redux';
import { RootState } from '../types';
import ServerStatusCard from '../components/dashboard/ServerStatusCard';
import MetricsChart from '../components/dashboard/MetricsChart';
import AlertsList from '../components/dashboard/AlertsList';
import OptimizationsList from '../components/dashboard/OptimizationsList';

const Dashboard: React.FC = () => {
    const servers = useSelector((state: RootState) => state.servers.items);
    const alerts = useSelector((state: RootState) => state.alerts.items);
    const metrics = useSelector((state: RootState) => state.metrics.current);
    const recommendations = useSelector(
        (state: RootState) => state.optimization.recommendations
    );

    return (
        <Grid container spacing={3}>
            <Grid item xs={12}>
                <Typography variant="h4" gutterBottom>
                    Панель мониторинга
                </Typography>
            </Grid>

            {/* Статус серверов */}
            <Grid item xs={12}>
                <Paper sx={{ p: 2 }}>
                    <Typography variant="h6" gutterBottom>
                        Статус серверов
                    </Typography>
                    <Grid container spacing={2}>
                        {servers.map((server) => (
                            <Grid item xs={12} sm={6} md={4} key={server.id}>
                                <ServerStatusCard
                                    server={server}
                                    metrics={metrics[server.id]}
                                />
                            </Grid>
                        ))}
                    </Grid>
                </Paper>
            </Grid>

            {/* Графики метрик */}
            <Grid item xs={12} md={8}>
                <Paper sx={{ p: 2 }}>
                    <Typography variant="h6" gutterBottom>
                        Метрики системы
                    </Typography>
                    <MetricsChart />
                </Paper>
            </Grid>

            {/* Активные алерты */}
            <Grid item xs={12} md={4}>
                <Paper sx={{ p: 2 }}>
                    <Typography variant="h6" gutterBottom>
                        Активные алерты
                    </Typography>
                    <AlertsList alerts={alerts} />
                </Paper>
            </Grid>

            {/* Рекомендации по оптимизации */}
            <Grid item xs={12}>
                <Paper sx={{ p: 2 }}>
                    <Typography variant="h6" gutterBottom>
                        Рекомендации по оптимизации
                    </Typography>
                    <OptimizationsList recommendations={recommendations} />
                </Paper>
            </Grid>
        </Grid>
    );
};

export default Dashboard; 