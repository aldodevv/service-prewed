package usecase

import (
	"context"
	"errors"
	"service-wedding/internal/domain"
)

type rsvpUsecase struct {
	rsvpRepo    domain.RSVPRepository
	guestRepo   domain.GuestRepository
	contextRepo domain.ContextRepository
}

func NewRSVPUsecase(rr domain.RSVPRepository, gr domain.GuestRepository, cr domain.ContextRepository) domain.RSVPUsecase {
	return &rsvpUsecase{
		rsvpRepo:    rr,
		guestRepo:   gr,
		contextRepo: cr,
	}
}

func (u *rsvpUsecase) GetByGuestID(ctx context.Context, guestID int64) (*domain.RSVP, error) {
	return u.rsvpRepo.GetByGuestID(ctx, guestID)
}

func (u *rsvpUsecase) GetAllByContextID(ctx context.Context, contextID int64) ([]domain.RSVPWithGuest, error) {
	// Verify context exists
	c, err := u.contextRepo.GetByID(ctx, contextID)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, errors.New("context not found")
	}

	return u.rsvpRepo.GetAllByContextID(ctx, contextID)
}

func (u *rsvpUsecase) SubmitRSVP(ctx context.Context, rsvp *domain.RSVP) error {
	if rsvp.GuestID <= 0 {
		return errors.New("invalid guest_id")
	}

	// Validate and normalize attendance status
	switch rsvp.Attendance {
	case "present", "attending":
		rsvp.Attendance = "attending"
	case "absent", "not_attending":
		rsvp.Attendance = "not_attending"
		rsvp.GuestCount = 0 // if not attending, count is 0
	case "undecided", "pending", "":
		rsvp.Attendance = "pending"
	default:
		return errors.New("invalid attendance status value")
	}

	if rsvp.Attendance == "attending" && rsvp.GuestCount <= 0 {
		rsvp.GuestCount = 1
	}

	// Upsert RSVP
	return u.rsvpRepo.Upsert(ctx, rsvp)
}
