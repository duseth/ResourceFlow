import React from 'react';
import WarningIcon from '@mui/icons-material/Warning';
import ErrorIcon from '@mui/icons-material/Error';
import InfoIcon from '@mui/icons-material/Info';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import { AlertType } from '../../store/types';

interface AlertIconProps {
  type: AlertType;
}

const AlertIcon: React.FC<AlertIconProps> = ({ type }) => {
  switch (type) {
    case 'warning':
      return <WarningIcon color="warning" />;
    case 'critical':
      return <ErrorIcon color="error" />;
    case 'info':
      return <InfoIcon color="info" />;
    case 'success':
      return <CheckCircleIcon color="success" />;
    default:
      return <InfoIcon color="info" />;
  }
};

export default AlertIcon; 