// Server types
export interface Server {
    id: string;
    name: string;
    host: string;
    port: number;
    status: ServerStatus;
    resourceUsage: ResourceUsage;
    createdAt: string;
    updatedAt: string;
}

export enum ServerStatus {
    ONLINE = 'ONLINE',
    OFFLINE = 'OFFLINE',
    MAINTENANCE = 'MAINTENANCE',
    WARNING = 'WARNING',
    ERROR = 'ERROR'
}

// Metric types
export interface Metric {
    id: string;
    serverId: string;
    timestamp: string;
    type: MetricType;
    value: number;
}

export enum MetricType {
    CPU_USAGE = 'CPU_USAGE',
    MEMORY_USAGE = 'MEMORY_USAGE',
    DISK_USAGE = 'DISK_USAGE',
    NETWORK_IN = 'NETWORK_IN',
    NETWORK_OUT = 'NETWORK_OUT'
}

export interface ResourceUsage {
    cpu: number;
    memory: number;
    disk: number;
    networkIn: number;
    networkOut: number;
}

// Alert types
export interface Alert {
    id: string;
    serverId: string;
    type: AlertType;
    message: string;
    status: AlertStatus;
    createdAt: string;
    resolvedAt?: string;
}

export enum AlertType {
    CPU = 'CPU',
    MEMORY = 'MEMORY',
    DISK = 'DISK',
    NETWORK = 'NETWORK',
    SERVER_STATUS = 'SERVER_STATUS'
}

export enum AlertStatus {
    ACTIVE = 'ACTIVE',
    RESOLVED = 'RESOLVED',
    ACKNOWLEDGED = 'ACKNOWLEDGED'
}

export interface AlertRule {
    id: string;
    name: string;
    type: AlertType;
    condition: AlertCondition;
    threshold: number;
    duration: number;
    enabled: boolean;
}

export interface AlertCondition {
    operator: 'GT' | 'LT' | 'GTE' | 'LTE' | 'EQ';
    value: number;
}

// Optimization types
export interface OptimizationRecommendation {
    id: string;
    serverId: string;
    type: OptimizationType;
    description: string;
    impact: OptimizationImpact;
    status: OptimizationStatus;
    createdAt: string;
    appliedAt?: string;
}

export enum OptimizationType {
    RESOURCE_ALLOCATION = 'RESOURCE_ALLOCATION',
    PERFORMANCE_TUNING = 'PERFORMANCE_TUNING',
    COST_OPTIMIZATION = 'COST_OPTIMIZATION',
    SECURITY = 'SECURITY'
}

export enum OptimizationImpact {
    HIGH = 'HIGH',
    MEDIUM = 'MEDIUM',
    LOW = 'LOW'
}

export enum OptimizationStatus {
    PENDING = 'PENDING',
    IN_PROGRESS = 'IN_PROGRESS',
    APPLIED = 'APPLIED',
    FAILED = 'FAILED'
}

// UI types
export interface Notification {
    id: string;
    type: 'success' | 'error' | 'warning' | 'info';
    message: string;
    duration?: number;
}

// Состояния для Redux
export interface RootState {
    servers: ServersState;
    metrics: MetricsState;
    alerts: AlertsState;
    optimization: OptimizationState;
    ui: UIState;
}

export interface ServersState {
    items: Server[];
    selectedServer: Server | null;
    loading: boolean;
    error: string | null;
}

export interface MetricsState {
    current: {
        [serverId: string]: ResourceUsage;
    };
    historical: {
        [serverId: string]: Metric[];
    };
    loading: boolean;
    error: string | null;
}

export interface AlertsState {
    items: Alert[];
    rules: AlertRule[];
    loading: boolean;
    error: string | null;
}

export interface OptimizationState {
    recommendations: OptimizationRecommendation[];
    loading: boolean;
    error: string | null;
}

export interface UIState {
    theme: 'light' | 'dark';
    sidebarOpen: boolean;
    notifications: Notification[];
}

// API интерфейсы
export interface ApiResponse<T> {
    data: T;
    status: number;
    message?: string;
}

// Фильтры и параметры запросов
export interface MetricsFilter {
    serverId?: string;
    types?: string[];
    from?: string;
    to?: string;
    interval?: string;
}

export interface ServerFilter {
    status?: string;
    tags?: string[];
    search?: string;
} 