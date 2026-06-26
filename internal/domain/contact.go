package domain

import (
	"context"
	"time"
)

type ContactMessage struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type ContactRepository interface {
	Create(ctx context.Context, msg *ContactMessage) error
	GetAll(ctx context.Context) ([]ContactMessage, error)
}
