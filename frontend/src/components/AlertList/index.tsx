import { FC } from 'react';
import { List, ListItem, ListItemIcon, ListItemText, Typography } from '@mui/material';
import { Alert, AlertType, AlertStatus } from '../../store/types';
import WarningIcon from '@mui/icons-material/Warning';
import ErrorIcon from '@mui/icons-material/Error';
import InfoIcon from '@mui/icons-material/Info';

interface AlertListProps {
  alerts: Alert[];
}

const getAlertIcon = (type: AlertType) => {
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

const getStatusColor = (status: AlertStatus): 'error' | 'warning' | 'success' | 'default' => {
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

const AlertList: FC<AlertListProps> = ({ alerts }) => {
  return (
    <List>
      {alerts.map((alert) => (
        <ListItem key={alert.id}>
          <ListItemIcon>{getAlertIcon(alert.type)}</ListItemIcon>
          <ListItemText 
            primary={alert.message}
            secondary={
              <>
                <Typography
                  component="span"
                  variant="body2"
                  color={getStatusColor(alert.status)}
                >
                  Status: {alert.status}
                </Typography>
                <br />
                <Typography component="span" variant="body2">
                  {alert.metricType.toUpperCase()}: {alert.metricValue} (Threshold: {alert.threshold})
                </Typography>
              </>
            }
          />
        </ListItem>
      ))}
    </List>
  );
};

export default AlertList; 