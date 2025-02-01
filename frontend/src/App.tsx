import React from 'react';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { Provider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import { theme } from './theme';
import { store } from './store';
import Layout from './components/Layout';
import Routes from './routes';

const App: React.FC = () => {
    return (
        <Provider store={store}>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                <BrowserRouter>
                    <Layout>
                        <Routes />
                    </Layout>
                </BrowserRouter>
            </ThemeProvider>
        </Provider>
    );
};

export default App; 