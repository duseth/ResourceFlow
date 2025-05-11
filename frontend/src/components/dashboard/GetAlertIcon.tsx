import {AlertType} from "../../store/types";
import WarningIcon from "@mui/icons-material/Warning";
import ErrorIcon from "@mui/icons-material/Error";

export const getAlertIcon = (type: AlertType) => {
    switch (type) {
        case 'warning':
            return <WarningIcon color="warning"/>;
        case 'critical':
            return <ErrorIcon color="error"/>;
        case 'error':
            return <ErrorIcon color="error"/>;
        default:
            return <WarningIcon/>;
    }
};