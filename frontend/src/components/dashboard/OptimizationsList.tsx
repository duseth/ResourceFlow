import React from 'react';
import {
    List,
    ListItem,
    ListItemText,
    ListItemIcon,
    Chip,
    Typography,
    Box,
    Button,
} from '@mui/material';
import {
    Memory as MemoryIcon,
    Storage as StorageIcon,
    Speed as SpeedIcon,
    Router as RouterIcon,
} from '@mui/icons-material';
import { OptimizationRecommendation } from '../../types';

interface OptimizationsListProps {
    recommendations: OptimizationRecommendation[];
}

const getTypeIcon = (type: OptimizationRecommendation['type']) => {
    switch (type) {
        case 'memory':
            return <MemoryIcon />;
        case 'disk':
            return <StorageIcon />;
        case 'cpu':
            return <SpeedIcon />;
        case 'network':
            return <RouterIcon />;
        default:
            return <SpeedIcon />;
    }
};

const getPriorityColor = (priority: OptimizationRecommendation['priority']) => {
    switch (priority) {
        case 'high':
            return 'error';
        case 'medium':
            return 'warning';
        case 'low':
            return 'info';
        default:
            return 'default';
    }
};

const getStatusColor = (status: OptimizationRecommendation['status']) => {
    switch (status) {
        case 'pending':
            return 'warning';
        case 'applied':
            return 'success';
        case 'rejected':
            return 'error';
        default:
            return 'default';
    }
};

const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString();
};

const OptimizationsList: React.FC<OptimizationsListProps> = ({ recommendations }) => {
    const sortedRecommendations = [...recommendations].sort((a, b) => {
        // Сортировка по приоритету
        const priorityOrder = { high: 0, medium: 1, low: 2 };
        const priorityDiff =
            priorityOrder[a.priority as keyof typeof priorityOrder] -
            priorityOrder[b.priority as keyof typeof priorityOrder];
        if (priorityDiff !== 0) return priorityDiff;

        // Если приоритеты равны, сортируем по дате создания
        return new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime();
    });

    return (
        <List>
            {sortedRecommendations.length === 0 ? (
                <ListItem>
                    <ListItemText primary="Нет рекомендаций по оптимизации" />
                </ListItem>
            ) : (
                sortedRecommendations.map((recommendation) => (
                    <ListItem
                        key={recommendation.id}
                        divider
                        secondaryAction={
                            recommendation.status === 'pending' && (
                                <Button variant="contained" color="primary" size="small">
                                    Применить
                                </Button>
                            )
                        }
                    >
                        <ListItemIcon>{getTypeIcon(recommendation.type)}</ListItemIcon>
                        <ListItemText
                            primary={
                                <Box display="flex" alignItems="center" gap={1}>
                                    <Typography variant="body1">
                                        {recommendation.description}
                                    </Typography>
                                    <Chip
                                        label={recommendation.priority}
                                        color={getPriorityColor(recommendation.priority)}
                                        size="small"
                                    />
                                    <Chip
                                        label={recommendation.status}
                                        color={getStatusColor(recommendation.status)}
                                        size="small"
                                    />
                                </Box>
                            }
                            secondary={
                                <>
                                    <Typography variant="body2" color="textSecondary">
                                        Сервер: {recommendation.serverId}
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary">
                                        Влияние: {recommendation.impact}
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary">
                                        Создано: {formatDate(recommendation.createdAt)}
                                    </Typography>
                                    {recommendation.appliedAt && (
                                        <Typography variant="body2" color="textSecondary">
                                            Применено: {formatDate(recommendation.appliedAt)}
                                        </Typography>
                                    )}
                                </>
                            }
                        />
                    </ListItem>
                ))
            )}
        </List>
    );
};

export default OptimizationsList; 