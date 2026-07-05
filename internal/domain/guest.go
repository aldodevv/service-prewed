package domain

import (
	"context"
	"time"
)

type Guest struct {
	ID        int64     `json:"id"`
	ContextID int64     `json:"context_id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	RSVP      *RSVP     `json:"rsvp,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GuestRepository interface {
	GetAllByContextID(ctx context.Context, contextID int64) ([]Guest, error)
	GetByID(ctx context.Context, contextID int64, guestID int64) (*Guest, error)
	GetBySlug(ctx context.Context, contextID int64, slug string) (*Guest, error)
	Create(ctx context.Context, guest *Guest) error
	Update(ctx context.Context, guest *Guest) error
	Delete(ctx context.Context, contextID int64, guestID int64) error
}
