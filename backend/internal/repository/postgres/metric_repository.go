package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/duseth/ResourceFlow/internal/domain/models"
)

type MetricRepository struct {
	db *sql.DB
}

func NewMetricRepository(db *sql.DB) *MetricRepository {
	return &MetricRepository{db: db}
}

func (r *MetricRepository) Store(ctx context.Context, metric *models.Metric) error {
	query := `
		INSERT INTO metrics (id, server_id, type, value, timestamp)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.ExecContext(ctx, query,
		metric.ID,
		metric.ServerID,
		metric.Type,
		metric.Value,
		metric.Timestamp,
	)

	if err != nil {
		return fmt.Errorf("failed to store metric: %v", err)
	}

	return nil
}

func (r *MetricRepository) GetLatest(ctx context.Context, serverID string, metricType string) (*models.Metric, error) {
	query := `
		SELECT id, server_id, type, value, timestamp
		FROM metrics
		WHERE server_id = $1 AND type = $2
		ORDER BY timestamp DESC
		LIMIT 1
	`

	var metric models.Metric
	err := r.db.QueryRowContext(ctx, query, serverID, metricType).Scan(
		&metric.ID,
		&metric.ServerID,
		&metric.Type,
		&metric.Value,
		&metric.Timestamp,
	)

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get latest metric: %v", err)
	}

	return &metric, nil
}

func (r *MetricRepository) GetRange(ctx context.Context, serverID string, from, to time.Time) ([]*models.Metric, error) {
	query := `
		SELECT id, server_id, type, value, timestamp
		FROM metrics
		WHERE server_id = $1 AND timestamp BETWEEN $2 AND $3
		ORDER BY timestamp ASC
	`

	rows, err := r.db.QueryContext(ctx, query, serverID, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get metrics range: %v", err)
	}
	defer rows.Close()

	var metrics []*models.Metric
	for rows.Next() {
		var metric models.Metric
		err := rows.Scan(
			&metric.ID,
			&metric.ServerID,
			&metric.Type,
			&metric.Value,
			&metric.Timestamp,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan metric: %v", err)
		}
		metrics = append(metrics, &metric)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating metrics: %v", err)
	}

	return metrics, nil
}

func (r *MetricRepository) GetAggregated(ctx context.Context, serverID string, period string) ([]*models.HistoricalData, error) {
	var interval string
	switch period {
	case models.HistoricalDataPeriodHour:
		interval = "1 hour"
	case models.HistoricalDataPeriodDay:
		interval = "1 day"
	case models.HistoricalDataPeriodWeek:
		interval = "1 week"
	case models.HistoricalDataPeriodMonth:
		interval = "1 month"
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	query := `
		WITH time_buckets AS (
			SELECT
				server_id,
				type,
				time_bucket($1::interval, timestamp) as bucket,
				MIN(value) as min_value,
				MAX(value) as max_value,
				AVG(value) as avg_value
			FROM metrics
			WHERE server_id = $2
			GROUP BY server_id, type, bucket
		)
		SELECT
			gen_random_uuid() as id,
			server_id,
			type,
			$3 as period,
			min_value,
			max_value,
			avg_value,
			bucket as timestamp
		FROM time_buckets
		ORDER BY bucket DESC
	`

	rows, err := r.db.QueryContext(ctx, query, interval, serverID, period)
	if err != nil {
		return nil, fmt.Errorf("failed to get aggregated metrics: %v", err)
	}
	defer rows.Close()

	var data []*models.HistoricalData
	for rows.Next() {
		var d models.HistoricalData
		err := rows.Scan(
			&d.ID,
			&d.ServerID,
			&d.MetricType,
			&d.Period,
			&d.MinValue,
			&d.MaxValue,
			&d.AvgValue,
			&d.Timestamp,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan historical data: %v", err)
		}
		data = append(data, &d)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating historical data: %v", err)
	}

	return data, nil
}
