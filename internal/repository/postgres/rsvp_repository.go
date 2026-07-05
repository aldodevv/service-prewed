package postgres

import (
	"context"
	"errors"
	"service-wedding/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type rsvpRepo struct {
	pool *pgxpool.Pool
}

func NewRSVPRepository(pool *pgxpool.Pool) domain.RSVPRepository {
	return &rsvpRepo{pool: pool}
}

func (r *rsvpRepo) GetByGuestID(ctx context.Context, guestID int64) (*domain.RSVP, error) {
	query := `SELECT id, guest_id, attendance, guest_count, message, created_at, updated_at FROM rsvps WHERE guest_id = $1`
	var rsvp domain.RSVP
	err := r.pool.QueryRow(ctx, query, guestID).Scan(
		&rsvp.ID, &rsvp.GuestID, &rsvp.Attendance, &rsvp.GuestCount, &rsvp.Message, &rsvp.CreatedAt, &rsvp.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &rsvp, nil
}

func (r *rsvpRepo) GetAllByContextID(ctx context.Context, contextID int64) ([]domain.RSVPWithGuest, error) {
	query := `
		SELECT r.id, r.guest_id, r.attendance, r.guest_count, r.message, r.created_at, r.updated_at, g.name
		FROM rsvps r
		JOIN guests g ON r.guest_id = g.id
		WHERE g.context_id = $1
		ORDER BY r.updated_at DESC
	`
	rows, err := r.pool.Query(ctx, query, contextID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rsvps []domain.RSVPWithGuest
	for rows.Next() {
		var item domain.RSVPWithGuest
		err := rows.Scan(
			&item.ID, &item.GuestID, &item.Attendance, &item.GuestCount, &item.Message, &item.CreatedAt, &item.UpdatedAt, &item.GuestName,
		)
		if err != nil {
			return nil, err
		}
		rsvps = append(rsvps, item)
	}
	return rsvps, nil
}

func (r *rsvpRepo) Upsert(ctx context.Context, rsvp *domain.RSVP) error {
	query := `
		INSERT INTO rsvps (guest_id, attendance, guest_count, message, updated_at)
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
		ON CONFLICT (guest_id) DO UPDATE
		SET attendance = EXCLUDED.attendance, guest_count = EXCLUDED.guest_count, message = EXCLUDED.message, updated_at = CURRENT_TIMESTAMP
		RETURNING id, created_at, updated_at
	`
	err := r.pool.QueryRow(ctx, query, rsvp.GuestID, rsvp.Attendance, rsvp.GuestCount, rsvp.Message).Scan(
		&rsvp.ID, &rsvp.CreatedAt, &rsvp.UpdatedAt,
	)
	return err
}
