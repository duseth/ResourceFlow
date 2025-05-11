import React from 'react';
import { List, ListItem, ListItemIcon, ListItemText, Typography, Box } from '@mui/material';
import { Optimization, OptimizationType, OptimizationStatus } from '../../store/types';
import SpeedIcon from '@mui/icons-material/Speed';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import ErrorIcon from '@mui/icons-material/Error';
import PendingIcon from '@mui/icons-material/Pending';
import BlockIcon from '@mui/icons-material/Block';

interface OptimizationsListProps {
  optimizations: Optimization[];
}

const getStatusIcon = (status: OptimizationStatus) => {
  switch (status) {
    case OptimizationStatus.PENDING:
      return <PendingIcon color="info" />;
    case OptimizationStatus.APPLIED:
      return <CheckCircleIcon color="success" />;
    case OptimizationStatus.REJECTED:
      return <BlockIcon color="error" />;
    case OptimizationStatus.IN_PROGRESS:
      return <SpeedIcon color="primary" />;
    case OptimizationStatus.FAILED:
      return <ErrorIcon color="error" />;
    default:
      return <PendingIcon />;
  }
};

const getStatusColor = (status: OptimizationStatus) => {
  switch (status) {
    case OptimizationStatus.PENDING:
      return 'info';
    case OptimizationStatus.APPLIED:
      return 'success';
    case OptimizationStatus.REJECTED:
      return 'error';
    case OptimizationStatus.IN_PROGRESS:
      return 'primary';
    case OptimizationStatus.FAILED:
      return 'error';
    default:
      return 'info';
  }
};

const getOptimizationTypeLabel = (type: OptimizationType) => {
  switch (type) {
    case OptimizationType.PERFORMANCE:
      return 'Performance';
    case OptimizationType.SCALING:
      return 'Scaling';
    case OptimizationType.COST:
      return 'Cost';
    default:
      return type;
  }
};

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleString();
};

const OptimizationsList: React.FC<OptimizationsListProps> = ({ optimizations }) => {
  return (
    <List>
      {optimizations.map((optimization) => (
        <ListItem key={optimization.id}>
          <ListItemIcon>{getStatusIcon(optimization.status)}</ListItemIcon>
          <ListItemText 
            primary={optimization.description}
            secondary={
              <Box component="span">
                <Typography
                  component="span"
                  variant="body2"
                  color={getStatusColor(optimization.status)}
                >
                  Status: {optimization.status.replace('_', ' ')}
                </Typography>
                <Box component="br" />
                <Typography component="span" variant="body2">
                  Type: {getOptimizationTypeLabel(optimization.type)}
                </Typography>
                <Box component="br" />
                <Typography component="span" variant="body2">
                  Impact: {optimization.impact}
                </Typography>
                <Box component="br" />
                <Typography component="span" variant="body2">
                  Created: {formatDate(optimization.createdAt)}
                  {optimization.appliedAt && ` | Applied: ${formatDate(optimization.appliedAt)}`}
                </Typography>
              </Box>
            }
          />
        </ListItem>
      ))}
    </List>
  );
};

export default OptimizationsList; 