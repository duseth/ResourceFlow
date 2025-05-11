import {FC} from 'react';
import {List, ListItem, ListItemIcon, ListItemText, Typography} from '@mui/material';
import {AlertListProps} from "./AlertListProps";
import {getAlertIcon} from "./GetAlertIcon";
import {getStatusColor} from "./GetStatusColor";

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