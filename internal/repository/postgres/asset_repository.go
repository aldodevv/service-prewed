package postgres

import (
	"context"
	"errors"
	"service-wedding/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type assetRepo struct {
	pool *pgxpool.Pool
}

func NewAssetRepository(pool *pgxpool.Pool) domain.AssetRepository {
	return &assetRepo{pool: pool}
}

func (r *assetRepo) GetAll(ctx context.Context) ([]domain.Asset, error) {
	query := `SELECT id, name, type, cloudinary_public_id, url, created_at, updated_at FROM assets ORDER BY id DESC`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []domain.Asset
	for rows.Next() {
		var a domain.Asset
		err := rows.Scan(&a.ID, &a.Name, &a.Type, &a.CloudinaryPublicID, &a.Url, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		assets = append(assets, a)
	}
	return assets, nil
}

func (r *assetRepo) GetByID(ctx context.Context, id int64) (*domain.Asset, error) {
	query := `SELECT id, name, type, cloudinary_public_id, url, created_at, updated_at FROM assets WHERE id = $1`
	var a domain.Asset
	err := r.pool.QueryRow(ctx, query, id).Scan(&a.ID, &a.Name, &a.Type, &a.CloudinaryPublicID, &a.Url, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *assetRepo) Create(ctx context.Context, a *domain.Asset) error {
	query := `INSERT INTO assets (name, type, cloudinary_public_id, url, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at`
	err := r.pool.QueryRow(ctx, query, a.Name, a.Type, a.CloudinaryPublicID, a.Url).Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt)
	return err
}

func (r *assetRepo) Update(ctx context.Context, a *domain.Asset) error {
	query := `UPDATE assets SET name = $1, type = $2, cloudinary_public_id = $3, url = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5`
	_, err := r.pool.Exec(ctx, query, a.Name, a.Type, a.CloudinaryPublicID, a.Url, a.ID)
	return err
}

func (r *assetRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM assets WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}
