import React from 'react';
import { Drawer, List, ListItemButton, ListItemIcon, ListItemText } from '@mui/material';
import { Link } from 'react-router-dom';
import DashboardIcon from '@mui/icons-material/Dashboard';
import StorageIcon from '@mui/icons-material/Storage';
import NotificationsIcon from '@mui/icons-material/Notifications';
import SpeedIcon from '@mui/icons-material/Speed';

const drawerWidth = 240;

const menuItems = [
  { text: 'Dashboard', icon: <DashboardIcon />, path: '/' },
  { text: 'Servers', icon: <StorageIcon />, path: '/servers' },
  { text: 'Alerts', icon: <NotificationsIcon />, path: '/alerts' },
  { text: 'Optimizations', icon: <SpeedIcon />, path: '/optimizations' },
];

const Sidebar: React.FC = () => {
  return (
    <Drawer
      variant="permanent"
      sx={{
        width: drawerWidth,
        flexShrink: 0,
        '& .MuiDrawer-paper': {
          width: drawerWidth,
          boxSizing: 'border-box',
          mt: 8, // header height
        },
      }}
    >
      <List>
        {menuItems.map((item) => (
          <ListItemButton key={item.text} component={Link} to={item.path}>
            <ListItemIcon>{item.icon}</ListItemIcon>
            <ListItemText primary={item.text} />
          </ListItemButton>
        ))}
      </List>
    </Drawer>
  );
};

export default Sidebar; 