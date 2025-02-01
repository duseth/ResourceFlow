package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/duseth/ResourceFlow/internal/domain/repository"
	"github.com/lib/pq"
)

type ServerRepository struct {
	db *sql.DB
}

func NewServerRepository(db *sql.DB) *ServerRepository {
	return &ServerRepository{db: db}
}

func (r *ServerRepository) Create(ctx context.Context, server *models.Server) error {
	tags, err := json.Marshal(server.Tags)
	if err != nil {
		return fmt.Errorf("failed to marshal tags: %v", err)
	}

	query := `
		INSERT INTO servers (id, name, host, port, status, tags, created_at, updated_at, last_check_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = r.db.ExecContext(ctx, query,
		server.ID,
		server.Name,
		server.Host,
		server.Port,
		server.Status,
		tags,
		server.CreatedAt,
		server.UpdatedAt,
		server.LastCheckAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create server: %v", err)
	}

	return nil
}

func (r *ServerRepository) Update(ctx context.Context, server *models.Server) error {
	tags, err := json.Marshal(server.Tags)
	if err != nil {
		return fmt.Errorf("failed to marshal tags: %v", err)
	}

	query := `
		UPDATE servers
		SET name = $1, host = $2, port = $3, status = $4, tags = $5, updated_at = $6, last_check_at = $7
		WHERE id = $8
	`

	result, err := r.db.ExecContext(ctx, query,
		server.Name,
		server.Host,
		server.Port,
		server.Status,
		tags,
		server.UpdatedAt,
		server.LastCheckAt,
		server.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update server: %v", err)
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

func (r *ServerRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM servers WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete server: %v", err)
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

func (r *ServerRepository) GetByID(ctx context.Context, id string) (*models.Server, error) {
	var server models.Server
	var tagsJSON []byte

	query := `
		SELECT id, name, host, port, status, tags, created_at, updated_at, last_check_at
		FROM servers
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&server.ID,
		&server.Name,
		&server.Host,
		&server.Port,
		&server.Status,
		&tagsJSON,
		&server.CreatedAt,
		&server.UpdatedAt,
		&server.LastCheckAt,
	)

	if err == sql.ErrNoRows {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get server: %v", err)
	}

	if err := json.Unmarshal(tagsJSON, &server.Tags); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tags: %v", err)
	}

	return &server, nil
}

func (r *ServerRepository) List(ctx context.Context, filter repository.ServerFilter) ([]*models.Server, error) {
	var args []interface{}
	query := `
		SELECT id, name, host, port, status, tags, created_at, updated_at, last_check_at
		FROM servers
		WHERE 1=1
	`

	if filter.Status != "" {
		args = append(args, filter.Status)
		query += fmt.Sprintf(" AND status = $%d", len(args))
	}

	if len(filter.Tags) > 0 {
		args = append(args, pq.Array(filter.Tags))
		query += fmt.Sprintf(" AND tags @> $%d", len(args))
	}

	if filter.Search != "" {
		args = append(args, "%"+filter.Search+"%")
		query += fmt.Sprintf(" AND (name ILIKE $%d OR host ILIKE $%d)", len(args), len(args))
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list servers: %v", err)
	}
	defer rows.Close()

	var servers []*models.Server
	for rows.Next() {
		var server models.Server
		var tagsJSON []byte

		err := rows.Scan(
			&server.ID,
			&server.Name,
			&server.Host,
			&server.Port,
			&server.Status,
			&tagsJSON,
			&server.CreatedAt,
			&server.UpdatedAt,
			&server.LastCheckAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan server: %v", err)
		}

		if err := json.Unmarshal(tagsJSON, &server.Tags); err != nil {
			return nil, fmt.Errorf("failed to unmarshal tags: %v", err)
		}

		servers = append(servers, &server)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating servers: %v", err)
	}

	return servers, nil
}
