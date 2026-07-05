package domain

import (
	"context"
	"time"
)

type RSVP struct {
	ID         int64     `json:"id"`
	GuestID    int64     `json:"guest_id"`
	Attendance string    `json:"attendance"` // 'attending', 'not_attending', 'pending'
	GuestCount int       `json:"guest_count"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RSVPWithGuest struct {
	RSVP
	GuestName string `json:"guest_name"`
}

type RSVPRepository interface {
	GetByGuestID(ctx context.Context, guestID int64) (*RSVP, error)
	GetAllByContextID(ctx context.Context, contextID int64) ([]RSVPWithGuest, error)
	Upsert(ctx context.Context, rsvp *RSVP) error
}

type RSVPUsecase interface {
	GetByGuestID(ctx context.Context, guestID int64) (*RSVP, error)
	GetAllByContextID(ctx context.Context, contextID int64) ([]RSVPWithGuest, error)
	SubmitRSVP(ctx context.Context, rsvp *RSVP) error
}
