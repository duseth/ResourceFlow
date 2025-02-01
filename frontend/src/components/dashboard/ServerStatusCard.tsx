import React from 'react';
import { Card, CardContent, Typography, Box, Chip, LinearProgress } from '@mui/material';
import { Server, ResourceUsage } from '../../types';

interface ServerStatusCardProps {
    server: Server;
    metrics: ResourceUsage;
}

const getStatusColor = (status: Server['status']) => {
    switch (status) {
        case 'online':
            return 'success';
        case 'warning':
            return 'warning';
        case 'error':
            return 'error';
        default:
            return 'default';
    }
};

const formatBytes = (bytes: number) => {
    if (bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`;
};

const ServerStatusCard: React.FC<ServerStatusCardProps> = ({ server, metrics }) => {
    return (
        <Card>
            <CardContent>
                <Box display="flex" justifyContent="space-between" alignItems="center" mb={2}>
                    <Typography variant="h6" component="h2">
                        {server.name}
                    </Typography>
                    <Chip
                        label={server.status}
                        color={getStatusColor(server.status)}
                        size="small"
                    />
                </Box>

                <Typography color="textSecondary" gutterBottom>
                    {server.host}:{server.port}
                </Typography>

                {metrics && (
                    <>
                        {/* CPU Usage */}
                        <Box mt={2}>
                            <Typography variant="body2" color="textSecondary">
                                CPU: {metrics.cpu.usage.toFixed(1)}%
                            </Typography>
                            <LinearProgress
                                variant="determinate"
                                value={metrics.cpu.usage}
                                color={metrics.cpu.usage > 80 ? 'error' : 'primary'}
                            />
                        </Box>

                        {/* Memory Usage */}
                        <Box mt={2}>
                            <Typography variant="body2" color="textSecondary">
                                Memory: {((metrics.memory.used / metrics.memory.total) * 100).toFixed(1)}%
                                ({formatBytes(metrics.memory.used)} / {formatBytes(metrics.memory.total)})
                            </Typography>
                            <LinearProgress
                                variant="determinate"
                                value={(metrics.memory.used / metrics.memory.total) * 100}
                                color={metrics.memory.used / metrics.memory.total > 0.8 ? 'error' : 'primary'}
                            />
                        </Box>

                        {/* Disk Usage */}
                        <Box mt={2}>
                            <Typography variant="body2" color="textSecondary">
                                Disk: {metrics.disk.utilization.toFixed(1)}%
                                ({formatBytes(metrics.disk.used)} / {formatBytes(metrics.disk.total)})
                            </Typography>
                            <LinearProgress
                                variant="determinate"
                                value={metrics.disk.utilization}
                                color={metrics.disk.utilization > 85 ? 'error' : 'primary'}
                            />
                        </Box>

                        {/* Network */}
                        <Box mt={2}>
                            <Typography variant="body2" color="textSecondary">
                                Network: ↓{formatBytes(metrics.network.bytesReceived)}/s 
                                ↑{formatBytes(metrics.network.bytesSent)}/s
                            </Typography>
                        </Box>
                    </>
                )}

                <Box mt={2} display="flex" flexWrap="wrap" gap={1}>
                    {server.tags.map((tag) => (
                        <Chip key={tag} label={tag} size="small" variant="outlined" />
                    ))}
                </Box>
            </CardContent>
        </Card>
    );
};

export default ServerStatusCard; 