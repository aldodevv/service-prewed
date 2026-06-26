package domain

import (
	"context"
	"time"
)

type Theme struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Slug        string                 `json:"slug"`
	Description string                 `json:"description"`
	Thumbnail   string                 `json:"thumbnail"`
	ThemeData   map[string]interface{} `json:"theme_data"`
	RenderHTML  string                 `json:"render_html"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type ThemeRepository interface {
	GetAll(ctx context.Context) ([]Theme, error)
	GetByID(ctx context.Context, id int64) (*Theme, error)
	GetBySlug(ctx context.Context, slug string) (*Theme, error)
	Create(ctx context.Context, theme *Theme) error
	Update(ctx context.Context, theme *Theme) error
	Delete(ctx context.Context, id int64) error
}
