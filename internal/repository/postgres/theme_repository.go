package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"service-wedding/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type themeRepo struct {
	pool *pgxpool.Pool
}

func NewThemeRepository(pool *pgxpool.Pool) domain.ThemeRepository {
	return &themeRepo{pool: pool}
}

func (r *themeRepo) GetAll(ctx context.Context) ([]domain.Theme, error) {
	query := `SELECT id, name, slug, description, thumbnail, theme_data, render_html, created_at, updated_at FROM themes ORDER BY id DESC`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var themes []domain.Theme
	for rows.Next() {
		var t domain.Theme
		var themeDataBytes []byte
		err := rows.Scan(
			&t.ID, &t.Name, &t.Slug, &t.Description, &t.Thumbnail, &themeDataBytes, &t.RenderHTML, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if len(themeDataBytes) > 0 {
			_ = json.Unmarshal(themeDataBytes, &t.ThemeData)
		}
		themes = append(themes, t)
	}
	return themes, nil
}

func (r *themeRepo) GetByID(ctx context.Context, id int64) (*domain.Theme, error) {
	query := `SELECT id, name, slug, description, thumbnail, theme_data, render_html, created_at, updated_at FROM themes WHERE id = $1`
	var t domain.Theme
	var themeDataBytes []byte
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&t.ID, &t.Name, &t.Slug, &t.Description, &t.Thumbnail, &themeDataBytes, &t.RenderHTML, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if len(themeDataBytes) > 0 {
		_ = json.Unmarshal(themeDataBytes, &t.ThemeData)
	}
	return &t, nil
}

func (r *themeRepo) GetBySlug(ctx context.Context, slug string) (*domain.Theme, error) {
	query := `SELECT id, name, slug, description, thumbnail, theme_data, render_html, created_at, updated_at FROM themes WHERE slug = $1`
	var t domain.Theme
	var themeDataBytes []byte
	err := r.pool.QueryRow(ctx, query, slug).Scan(
		&t.ID, &t.Name, &t.Slug, &t.Description, &t.Thumbnail, &themeDataBytes, &t.RenderHTML, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if len(themeDataBytes) > 0 {
		_ = json.Unmarshal(themeDataBytes, &t.ThemeData)
	}
	return &t, nil
}

func (r *themeRepo) Create(ctx context.Context, t *domain.Theme) error {
	themeDataBytes, err := json.Marshal(t.ThemeData)
	if err != nil {
		return err
	}

	query := `INSERT INTO themes (name, slug, description, thumbnail, theme_data, render_html, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at`
	err = r.pool.QueryRow(ctx, query, t.Name, t.Slug, t.Description, t.Thumbnail, themeDataBytes, t.RenderHTML).
		Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
	return err
}

func (r *themeRepo) Update(ctx context.Context, t *domain.Theme) error {
	themeDataBytes, err := json.Marshal(t.ThemeData)
	if err != nil {
		return err
	}

	query := `UPDATE themes SET name = $1, slug = $2, description = $3, thumbnail = $4, theme_data = $5, render_html = $6, updated_at = CURRENT_TIMESTAMP WHERE id = $7`
	_, err = r.pool.Exec(ctx, query, t.Name, t.Slug, t.Description, t.Thumbnail, themeDataBytes, t.RenderHTML, t.ID)
	return err
}

func (r *themeRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM themes WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}
