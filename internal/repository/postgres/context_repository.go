package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"service-wedding/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type contextRepo struct {
	pool *pgxpool.Pool
}

func NewContextRepository(pool *pgxpool.Pool) domain.ContextRepository {
	return &contextRepo{pool: pool}
}

func (r *contextRepo) GetAll(ctx context.Context) ([]domain.Context, error) {
	query := `SELECT id, name, slug, COALESCE(theme_id, 0) AS theme_id, render_html, content_json, created_at, updated_at FROM contexts ORDER BY id DESC`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contexts []domain.Context
	for rows.Next() {
		var c domain.Context
		var contentJsonBytes []byte
		err := rows.Scan(
			&c.ID, &c.Name, &c.Slug, &c.ThemeID, &c.RenderHTML, &contentJsonBytes, &c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if len(contentJsonBytes) > 0 {
			_ = json.Unmarshal(contentJsonBytes, &c.ContentJSON)
		}
		contexts = append(contexts, c)
	}
	return contexts, nil
}

func (r *contextRepo) GetByID(ctx context.Context, id int64) (*domain.Context, error) {
	query := `SELECT id, name, slug, COALESCE(theme_id, 0) AS theme_id, render_html, content_json, created_at, updated_at FROM contexts WHERE id = $1`
	var c domain.Context
	var contentJsonBytes []byte
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&c.ID, &c.Name, &c.Slug, &c.ThemeID, &c.RenderHTML, &contentJsonBytes, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if len(contentJsonBytes) > 0 {
		_ = json.Unmarshal(contentJsonBytes, &c.ContentJSON)
	}
	return &c, nil
}

func (r *contextRepo) GetBySlug(ctx context.Context, slug string) (*domain.Context, error) {
	query := `SELECT id, name, slug, COALESCE(theme_id, 0) AS theme_id, render_html, content_json, created_at, updated_at FROM contexts WHERE slug = $1`
	var c domain.Context
	var contentJsonBytes []byte
	err := r.pool.QueryRow(ctx, query, slug).Scan(
		&c.ID, &c.Name, &c.Slug, &c.ThemeID, &c.RenderHTML, &contentJsonBytes, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if len(contentJsonBytes) > 0 {
		_ = json.Unmarshal(contentJsonBytes, &c.ContentJSON)
	}
	return &c, nil
}

func (r *contextRepo) Create(ctx context.Context, c *domain.Context) error {
	contentJsonBytes, err := json.Marshal(c.ContentJSON)
	if err != nil {
		return err
	}

	query := `INSERT INTO contexts (name, slug, theme_id, render_html, content_json, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at`
	err = r.pool.QueryRow(ctx, query, c.Name, c.Slug, c.ThemeID, c.RenderHTML, contentJsonBytes).
		Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt)
	return err
}

func (r *contextRepo) Update(ctx context.Context, c *domain.Context) error {
	contentJsonBytes, err := json.Marshal(c.ContentJSON)
	if err != nil {
		return err
	}

	query := `UPDATE contexts SET name = $1, slug = $2, theme_id = $3, render_html = $4, content_json = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $6`
	_, err = r.pool.Exec(ctx, query, c.Name, c.Slug, c.ThemeID, c.RenderHTML, contentJsonBytes, c.ID)
	return err
}

func (r *contextRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM contexts WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}
