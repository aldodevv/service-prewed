package domain

import (
	"context"
	"time"
)

type Context struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Slug        string                 `json:"slug"`
	ThemeID     int64                  `json:"theme_id"`
	RenderHTML  string                 `json:"render_html"`
	ContentJSON map[string]interface{} `json:"content_json"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type ContextRepository interface {
	GetAll(ctx context.Context) ([]Context, error)
	GetByID(ctx context.Context, id int64) (*Context, error)
	GetBySlug(ctx context.Context, slug string) (*Context, error)
	Create(ctx context.Context, c *Context) error
	Update(ctx context.Context, c *Context) error
	Delete(ctx context.Context, id int64) error
}
