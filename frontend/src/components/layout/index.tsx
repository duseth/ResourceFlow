import React from 'react';
import { Box, styled } from '@mui/material';
import Header from '../Header';
import Sidebar from '../Sidebar';

interface LayoutProps {
  children: React.ReactNode;
}

const Main = styled(Box)(({ theme }) => ({
  flexGrow: 1,
  padding: theme.spacing(3),
  marginLeft: theme.custom.sidebar.width,
  marginTop: theme.custom.header.height,
  minHeight: `calc(100vh - ${theme.custom.header.height}px)`,
  background: theme.palette.background.default,
}));

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <Box sx={{ display: 'flex' }}>
      <Header />
      <Sidebar />
      <Main>{children}</Main>
    </Box>
  );
};

export default Layout; 