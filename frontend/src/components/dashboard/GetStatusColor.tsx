import {AlertStatus} from "../../store/types";

export const getStatusColor = (status: AlertStatus): 'error' | 'warning' | 'success' | 'default' => {
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