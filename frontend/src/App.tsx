import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Layout from './components/Layout';
import Dashboard from './pages/Dashboard';
import Servers from './pages/Servers';
import Alerts from './pages/Alerts';
import Optimizations from './pages/Optimizations';

const App: React.FC = () => {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/servers" element={<Servers />} />
          <Route path="/alerts" element={<Alerts />} />
          <Route path="/optimizations" element={<Optimizations />} />
        </Routes>
      </Layout>
    </Router>
  );
};

export default App; 