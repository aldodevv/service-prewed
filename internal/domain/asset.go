package domain

import (
	"context"
	"time"
)

type Asset struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	Type               string    `json:"type"` // e.g. "image", "video", "audio", "font"
	CloudinaryPublicID string    `json:"cloudinary_public_id"`
	Url                string    `json:"url"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type AssetRepository interface {
	GetAll(ctx context.Context) ([]Asset, error)
	GetByID(ctx context.Context, id int64) (*Asset, error)
	Create(ctx context.Context, asset *Asset) error
	Update(ctx context.Context, asset *Asset) error
	Delete(ctx context.Context, id int64) error
}
