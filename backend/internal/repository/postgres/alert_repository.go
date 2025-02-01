package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/duseth/ResourceFlow/internal/domain/repository"
)

type AlertRepository struct {
	db *sql.DB
}

func NewAlertRepository(db *sql.DB) *AlertRepository {
	return &AlertRepository{db: db}
}

func (r *AlertRepository) Create(ctx context.Context, alert *models.Alert) error {
	query := `
		INSERT INTO alerts (id, server_id, rule_id, status, message, created_at, resolved_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		alert.ID,
		alert.ServerID,
		alert.RuleID,
		alert.Status,
		alert.Message,
		alert.CreatedAt,
		alert.ResolvedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create alert: %v", err)
	}

	return nil
}

func (r *AlertRepository) Update(ctx context.Context, alert *models.Alert) error {
	query := `
		UPDATE alerts
		SET status = $1, resolved_at = $2
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query,
		alert.Status,
		alert.ResolvedAt,
		alert.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update alert: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rows == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *AlertRepository) GetByID(ctx context.Context, id string) (*models.Alert, error) {
	query := `
		SELECT id, server_id, rule_id, status, message, created_at, resolved_at
		FROM alerts
		WHERE id = $1
	`

	var alert models.Alert
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&alert.ID,
		&alert.ServerID,
		&alert.RuleID,
		&alert.Status,
		&alert.Message,
		&alert.CreatedAt,
		&alert.ResolvedAt,
	)

	if err == sql.ErrNoRows {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get alert: %v", err)
	}

	return &alert, nil
}

func (r *AlertRepository) GetActive(ctx context.Context) ([]*models.Alert, error) {
	query := `
		SELECT id, server_id, rule_id, status, message, created_at, resolved_at
		FROM alerts
		WHERE status = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, models.AlertStatusActive)
	if err != nil {
		return nil, fmt.Errorf("failed to get active alerts: %v", err)
	}
	defer rows.Close()

	var alerts []*models.Alert
	for rows.Next() {
		var alert models.Alert
		err := rows.Scan(
			&alert.ID,
			&alert.ServerID,
			&alert.RuleID,
			&alert.Status,
			&alert.Message,
			&alert.CreatedAt,
			&alert.ResolvedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan alert: %v", err)
		}
		alerts = append(alerts, &alert)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating alerts: %v", err)
	}

	return alerts, nil
}

func (r *AlertRepository) GetByServer(ctx context.Context, serverID string) ([]*models.Alert, error) {
	query := `
		SELECT id, server_id, rule_id, status, message, created_at, resolved_at
		FROM alerts
		WHERE server_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, serverID)
	if err != nil {
		return nil, fmt.Errorf("failed to get server alerts: %v", err)
	}
	defer rows.Close()

	var alerts []*models.Alert
	for rows.Next() {
		var alert models.Alert
		err := rows.Scan(
			&alert.ID,
			&alert.ServerID,
			&alert.RuleID,
			&alert.Status,
			&alert.Message,
			&alert.CreatedAt,
			&alert.ResolvedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan alert: %v", err)
		}
		alerts = append(alerts, &alert)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating alerts: %v", err)
	}

	return alerts, nil
}

// Методы для работы с правилами алертов

func (r *AlertRepository) GetRules(ctx context.Context, serverID string) ([]*models.AlertRule, error) {
	query := `
		SELECT id, server_id, metric_type, condition, threshold, duration
		FROM alert_rules
		WHERE server_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, serverID)
	if err != nil {
		return nil, fmt.Errorf("failed to get alert rules: %v", err)
	}
	defer rows.Close()

	var rules []*models.AlertRule
	for rows.Next() {
		var rule models.AlertRule
		err := rows.Scan(
			&rule.ID,
			&rule.ServerID,
			&rule.MetricType,
			&rule.Condition,
			&rule.Threshold,
			&rule.Duration,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan alert rule: %v", err)
		}
		rules = append(rules, &rule)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating alert rules: %v", err)
	}

	return rules, nil
}

func (r *AlertRepository) CreateRule(ctx context.Context, rule *models.AlertRule) error {
	query := `
		INSERT INTO alert_rules (id, server_id, metric_type, condition, threshold, duration)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		rule.ID,
		rule.ServerID,
		rule.MetricType,
		rule.Condition,
		rule.Threshold,
		rule.Duration,
	)

	if err != nil {
		return fmt.Errorf("failed to create alert rule: %v", err)
	}

	return nil
}

func (r *AlertRepository) UpdateRule(ctx context.Context, rule *models.AlertRule) error {
	query := `
		UPDATE alert_rules
		SET metric_type = $1, condition = $2, threshold = $3, duration = $4
		WHERE id = $5 AND server_id = $6
	`

	result, err := r.db.ExecContext(ctx, query,
		rule.MetricType,
		rule.Condition,
		rule.Threshold,
		rule.Duration,
		rule.ID,
		rule.ServerID,
	)

	if err != nil {
		return fmt.Errorf("failed to update alert rule: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rows == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *AlertRepository) DeleteRule(ctx context.Context, id string) error {
	query := `DELETE FROM alert_rules WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete alert rule: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rows == 0 {
		return repository.ErrNotFound
	}

	return nil
}
