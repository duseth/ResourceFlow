package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/duseth/ResourceFlow/internal/domain/repository"
)

type OptimizationRepository struct {
	db *sql.DB
}

func NewOptimizationRepository(db *sql.DB) *OptimizationRepository {
	return &OptimizationRepository{db: db}
}

func (r *OptimizationRepository) CreateRecommendation(ctx context.Context, rec *models.OptimizationRecommendation) error {
	query := `
		INSERT INTO optimization_recommendations (
			id, server_id, type, description, impact, status, created_at, applied_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(ctx, query,
		rec.ID,
		rec.ServerID,
		rec.Type,
		rec.Description,
		rec.Impact,
		rec.Status,
		rec.CreatedAt,
		rec.AppliedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create optimization recommendation: %v", err)
	}

	return nil
}

func (r *OptimizationRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `
		UPDATE optimization_recommendations
		SET status = $1, applied_at = $2
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, status, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update optimization status: %v", err)
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

func (r *OptimizationRepository) GetPending(ctx context.Context) ([]*models.OptimizationRecommendation, error) {
	query := `
		SELECT id, server_id, type, description, impact, status, created_at, applied_at
		FROM optimization_recommendations
		WHERE status = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, models.OptimizationStatusPending)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending optimizations: %v", err)
	}
	defer rows.Close()

	var recommendations []*models.OptimizationRecommendation
	for rows.Next() {
		var rec models.OptimizationRecommendation
		err := rows.Scan(
			&rec.ID,
			&rec.ServerID,
			&rec.Type,
			&rec.Description,
			&rec.Impact,
			&rec.Status,
			&rec.CreatedAt,
			&rec.AppliedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan optimization: %v", err)
		}
		recommendations = append(recommendations, &rec)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating optimizations: %v", err)
	}

	return recommendations, nil
}

func (r *OptimizationRepository) GetByServer(ctx context.Context, serverID string) ([]*models.OptimizationRecommendation, error) {
	query := `
		SELECT id, server_id, type, description, impact, status, created_at, applied_at
		FROM optimization_recommendations
		WHERE server_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, serverID)
	if err != nil {
		return nil, fmt.Errorf("failed to get server optimizations: %v", err)
	}
	defer rows.Close()

	var recommendations []*models.OptimizationRecommendation
	for rows.Next() {
		var rec models.OptimizationRecommendation
		err := rows.Scan(
			&rec.ID,
			&rec.ServerID,
			&rec.Type,
			&rec.Description,
			&rec.Impact,
			&rec.Status,
			&rec.CreatedAt,
			&rec.AppliedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan optimization: %v", err)
		}
		recommendations = append(recommendations, &rec)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating optimizations: %v", err)
	}

	return recommendations, nil
}
