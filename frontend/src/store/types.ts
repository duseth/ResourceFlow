export type AlertType = 'warning' | 'critical' | 'error' | 'success' | 'info';
export type AlertStatus = 'active' | 'acknowledged' | 'resolved';
export type MetricType = 'cpu' | 'memory' | 'disk' | 'network';

export interface Alert {
  id: string;
  type: AlertType;
  message: string;
  status: AlertStatus;
  serverId: string;
  createdAt: string;
  resolvedAt?: string;
  acknowledgedAt?: string;
  metricType: MetricType;
  metricValue: number;
  threshold: number;
}

export enum ServerStatus {
  ONLINE = 'online',
  OFFLINE = 'offline',
  MAINTENANCE = 'maintenance'
}

export interface Server {
  id: string;
  name: string;
  status: ServerStatus;
  ip: string;
  cpu: number;
  memory: number;
  disk: number;
  lastSeen: string;
  maintenanceScheduled?: string;
  location?: string;
  tags?: string[];
  description?: string;
}

export enum OptimizationType {
  PERFORMANCE = 'performance',
  SCALING = 'scaling',
  COST = 'cost'
}

export enum OptimizationStatus {
  PENDING = 'pending',
  APPLIED = 'applied',
  REJECTED = 'rejected',
  IN_PROGRESS = 'in_progress',
  FAILED = 'failed'
}

export interface Optimization {
  id: string;
  type: OptimizationType;
  description: string;
  status: OptimizationStatus;
  serverId: string;
  impact: string;
  createdAt: string;
  appliedAt?: string;
} 