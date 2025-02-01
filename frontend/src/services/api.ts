import axios from 'axios';
import {
    Server,
    Metric,
    Alert,
    AlertRule,
    OptimizationRecommendation,
    ApiResponse,
    ServerFilter,
    MetricsFilter,
} from '../types';

const BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

const axiosInstance = axios.create({
    baseURL: BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

export class Api {
    // Servers
    static getServers = () => axiosInstance.get<Server[]>('/servers');
    static getServer = (id: string) => axiosInstance.get<Server>(`/servers/${id}`);
    static createServer = (data: Omit<Server, 'id'>) => axiosInstance.post<Server>('/servers', data);
    static updateServer = (id: string, data: Server) => axiosInstance.put<Server>(`/servers/${id}`, data);
    static deleteServer = (id: string) => axiosInstance.delete(`/servers/${id}`);

    // Metrics
    static getLiveMetrics = (serverId: string) => axiosInstance.get<Metric[]>(`/servers/${serverId}/metrics/live`);
    static getHistoricalMetrics = (serverId: string, startDate: string, endDate: string) =>
        axiosInstance.get<Metric[]>(`/servers/${serverId}/metrics/historical`, {
            params: { startDate, endDate },
        });

    // Alerts
    static getAlerts = () => axiosInstance.get<Alert[]>('/alerts');
    static getAlertRules = () => axiosInstance.get<AlertRule[]>('/alert-rules');
    static createAlertRule = (data: Omit<AlertRule, 'id'>) => axiosInstance.post<AlertRule>('/alert-rules', data);
    static updateAlertRule = (id: string, data: AlertRule) => axiosInstance.put<AlertRule>(`/alert-rules/${id}`, data);
    static deleteAlertRule = (id: string) => axiosInstance.delete(`/alert-rules/${id}`);

    // Optimization
    static getOptimizationRecommendations = (serverId: string) =>
        axiosInstance.get<OptimizationRecommendation[]>(`/servers/${serverId}/recommendations`);
    static applyOptimizationRecommendation = (serverId: string, recommendationId: string) =>
        axiosInstance.post<OptimizationRecommendation>(
            `/servers/${serverId}/recommendations/${recommendationId}/apply`
        );
}

export const api = Api; 