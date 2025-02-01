-- Создание таблицы servers
CREATE TABLE IF NOT EXISTS servers (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    host VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL,
    status VARCHAR(50) NOT NULL,
    tags JSONB NOT NULL DEFAULT '[]',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    last_check_at TIMESTAMP NOT NULL
);

-- Создание таблицы metrics
CREATE TABLE IF NOT EXISTS metrics (
    id UUID PRIMARY KEY,
    server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    value DOUBLE PRECISION NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

-- Создание таблицы alert_rules
CREATE TABLE IF NOT EXISTS alert_rules (
    id UUID PRIMARY KEY,
    server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
    metric_type VARCHAR(50) NOT NULL,
    condition VARCHAR(50) NOT NULL,
    threshold DOUBLE PRECISION NOT NULL,
    duration INTEGER NOT NULL
);

-- Создание таблицы alerts
CREATE TABLE IF NOT EXISTS alerts (
    id UUID PRIMARY KEY,
    server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
    rule_id UUID REFERENCES alert_rules(id) ON DELETE SET NULL,
    status VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    resolved_at TIMESTAMP
);

-- Создание таблицы optimization_recommendations
CREATE TABLE IF NOT EXISTS optimization_recommendations (
    id UUID PRIMARY KEY,
    server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    impact TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    applied_at TIMESTAMP
);

-- Создание индексов
CREATE INDEX idx_metrics_server_id ON metrics(server_id);
CREATE INDEX idx_metrics_timestamp ON metrics(timestamp);
CREATE INDEX idx_metrics_type ON metrics(type);
CREATE INDEX idx_alerts_server_id ON alerts(server_id);
CREATE INDEX idx_alerts_status ON alerts(status);
CREATE INDEX idx_alerts_created_at ON alerts(created_at);
CREATE INDEX idx_alert_rules_server_id ON alert_rules(server_id);
CREATE INDEX idx_optimizations_server_id ON optimization_recommendations(server_id);
CREATE INDEX idx_optimizations_status ON optimization_recommendations(status);
CREATE INDEX idx_optimizations_created_at ON optimization_recommendations(created_at);

-- Добавляем расширение для TimescaleDB
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;

-- Создаем гипертаблицу для метрик
SELECT create_hypertable('metrics', 'timestamp', 
    chunk_time_interval => INTERVAL '1 day',
    if_not_exists => TRUE
); 