import React from 'react';
import {Box, Card, CardContent, LinearProgress, Typography} from '@mui/material';
import {Server, ServerStatus} from '../../store/types';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import ErrorIcon from '@mui/icons-material/Error';
import WarningIcon from '@mui/icons-material/Warning';

interface ServerStatusCardProps {
    server: Server;
}

const getStatusIcon = (status: ServerStatus) => {
    switch (status) {
        case ServerStatus.ONLINE:
            return <CheckCircleIcon color="success" />;
        case ServerStatus.OFFLINE:
            return <ErrorIcon color="error" />;
        case ServerStatus.MAINTENANCE:
            return <WarningIcon color="warning" />;
        default:
            return <ErrorIcon color="error" />;
    }
};

const getStatusColor = (status: ServerStatus) => {
    switch (status) {
        case ServerStatus.ONLINE:
            return 'success.main';
        case ServerStatus.OFFLINE:
            return 'error.main';
        case ServerStatus.MAINTENANCE:
            return 'warning.main';
        default:
            return 'error.main';
    }
};

const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleString();
};

const ServerStatusCard: React.FC<ServerStatusCardProps> = ({ server }) => {
    return (
        <Card>
            <CardContent>
                <Box display="flex" alignItems="center" mb={2}>
                    <Box mr={1}>{getStatusIcon(server.status)}</Box>
                    <Typography variant="h6" component="div">
                        {server.name}
                    </Typography>
                </Box>

                <Typography color="textSecondary" gutterBottom>
                    IP: {server.ip}
                </Typography>

                <Box my={2}>
                    <Typography variant="body2" gutterBottom>
                        CPU Usage
                    </Typography>
                    <LinearProgress
                        variant="determinate"
                        value={server.cpu}
                        color={server.cpu > 80 ? "error" : server.cpu > 60 ? "warning" : "primary"}
                    />
                    <Typography variant="body2" color="textSecondary">
                        {server.cpu}%
                    </Typography>
                </Box>

                <Box my={2}>
                    <Typography variant="body2" gutterBottom>
                        Memory Usage
                    </Typography>
                    <LinearProgress
                        variant="determinate"
                        value={server.memory}
                        color={server.memory > 80 ? "error" : server.memory > 60 ? "warning" : "primary"}
                    />
                    <Typography variant="body2" color="textSecondary">
                        {server.memory}%
                    </Typography>
                </Box>

                <Box my={2}>
                    <Typography variant="body2" gutterBottom>
                        Disk Usage
                    </Typography>
                    <LinearProgress
                        variant="determinate"
                        value={server.disk}
                        color={server.disk > 85 ? "error" : server.disk > 70 ? "warning" : "primary"}
                    />
                    <Typography variant="body2" color="textSecondary">
                        {server.disk}%
                    </Typography>
                </Box>

                <Typography variant="body2" color="textSecondary">
                    Last Seen: {formatDate(server.lastSeen)}
                </Typography>

                {server.maintenanceScheduled && (
                    <Typography variant="body2" color="warning.main">
                        Maintenance Scheduled: {formatDate(server.maintenanceScheduled)}
                    </Typography>
                )}
            </CardContent>
        </Card>
    );
};

export default ServerStatusCard;