package postgres

import (
	"context"
	"errors"
	"service-wedding/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type guestRepo struct {
	pool *pgxpool.Pool
}

func NewGuestRepository(pool *pgxpool.Pool) domain.GuestRepository {
	return &guestRepo{pool: pool}
}

func (r *guestRepo) GetAllByContextID(ctx context.Context, contextID int64) ([]domain.Guest, error) {
	query := `SELECT id, context_id, name, slug, created_at, updated_at FROM guests WHERE context_id = $1 ORDER BY id DESC`
	rows, err := r.pool.Query(ctx, query, contextID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guests []domain.Guest
	for rows.Next() {
		var g domain.Guest
		err := rows.Scan(&g.ID, &g.ContextID, &g.Name, &g.Slug, &g.CreatedAt, &g.UpdatedAt)
		if err != nil {
			return nil, err
		}
		guests = append(guests, g)
	}
	return guests, nil
}

func (r *guestRepo) GetByID(ctx context.Context, contextID int64, guestID int64) (*domain.Guest, error) {
	query := `SELECT id, context_id, name, slug, created_at, updated_at FROM guests WHERE context_id = $1 AND id = $2`
	var g domain.Guest
	err := r.pool.QueryRow(ctx, query, contextID, guestID).Scan(&g.ID, &g.ContextID, &g.Name, &g.Slug, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &g, nil
}

func (r *guestRepo) GetBySlug(ctx context.Context, contextID int64, slug string) (*domain.Guest, error) {
	query := `SELECT id, context_id, name, slug, created_at, updated_at FROM guests WHERE context_id = $1 AND slug = $2`
	var g domain.Guest
	err := r.pool.QueryRow(ctx, query, contextID, slug).Scan(&g.ID, &g.ContextID, &g.Name, &g.Slug, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &g, nil
}

func (r *guestRepo) Create(ctx context.Context, g *domain.Guest) error {
	query := `INSERT INTO guests (context_id, name, slug, created_at, updated_at) 
	          VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at`
	err := r.pool.QueryRow(ctx, query, g.ContextID, g.Name, g.Slug).Scan(&g.ID, &g.CreatedAt, &g.UpdatedAt)
	return err
}

func (r *guestRepo) Update(ctx context.Context, g *domain.Guest) error {
	query := `UPDATE guests SET name = $1, slug = $2, updated_at = CURRENT_TIMESTAMP WHERE context_id = $3 AND id = $4`
	_, err := r.pool.Exec(ctx, query, g.Name, g.Slug, g.ContextID, g.ID)
	return err
}

func (r *guestRepo) Delete(ctx context.Context, contextID int64, guestID int64) error {
	query := `DELETE FROM guests WHERE context_id = $1 AND id = $2`
	_, err := r.pool.Exec(ctx, query, contextID, guestID)
	return err
}
