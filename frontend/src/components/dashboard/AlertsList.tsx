import React from 'react';
import {
    List,
    ListItem,
    ListItemText,
    ListItemIcon,
    Chip,
    Typography,
    Box,
} from '@mui/material';
import {
    Warning as WarningIcon,
    Error as ErrorIcon,
    CheckCircle as CheckCircleIcon,
} from '@mui/icons-material';
import { Alert } from '../../types';

interface AlertsListProps {
    alerts: Alert[];
}

const getAlertIcon = (type: Alert['type']) => {
    switch (type) {
        case 'warning':
            return <WarningIcon color="warning" />;
        case 'critical':
            return <ErrorIcon color="error" />;
        case 'error':
            return <ErrorIcon color="error" />;
        default:
            return <WarningIcon />;
    }
};

const getStatusColor = (status: Alert['status']) => {
    switch (status) {
        case 'active':
            return 'error';
        case 'acknowledged':
            return 'warning';
        case 'resolved':
            return 'success';
        default:
            return 'default';
    }
};

const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString();
};

const AlertsList: React.FC<AlertsListProps> = ({ alerts }) => {
    const sortedAlerts = [...alerts].sort(
        (a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
    );

    return (
        <List>
            {sortedAlerts.length === 0 ? (
                <ListItem>
                    <ListItemIcon>
                        <CheckCircleIcon color="success" />
                    </ListItemIcon>
                    <ListItemText primary="Нет активных алертов" />
                </ListItem>
            ) : (
                sortedAlerts.map((alert) => (
                    <ListItem key={alert.id} divider>
                        <ListItemIcon>{getAlertIcon(alert.type)}</ListItemIcon>
                        <ListItemText
                            primary={
                                <Box display="flex" alignItems="center" gap={1}>
                                    <Typography variant="body1">{alert.message}</Typography>
                                    <Chip
                                        label={alert.status}
                                        color={getStatusColor(alert.status)}
                                        size="small"
                                    />
                                </Box>
                            }
                            secondary={
                                <>
                                    <Typography variant="body2" color="textSecondary">
                                        Сервер: {alert.serverId}
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary">
                                        Создан: {formatDate(alert.createdAt)}
                                    </Typography>
                                    {alert.resolvedAt && (
                                        <Typography variant="body2" color="textSecondary">
                                            Решен: {formatDate(alert.resolvedAt)}
                                        </Typography>
                                    )}
                                    <Typography variant="body2" color="textSecondary">
                                        {alert.metricType}: {alert.metricValue} (порог: {alert.threshold})
                                    </Typography>
                                </>
                            }
                        />
                    </ListItem>
                ))
            )}
        </List>
    );
};

export default AlertsList; 