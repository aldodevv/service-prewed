package usecase

import (
	"context"
	"errors"
	"service-wedding/internal/domain"
)

type GuestUsecase interface {
	GetAllByContextID(ctx context.Context, contextID int64) ([]domain.Guest, error)
	GetByID(ctx context.Context, contextID int64, guestID int64) (*domain.Guest, error)
	Create(ctx context.Context, guest *domain.Guest) error
	Update(ctx context.Context, contextID int64, guestID int64, name, slug string) error
	Delete(ctx context.Context, contextID int64, guestID int64) error
}

type guestUsecase struct {
	guestRepo   domain.GuestRepository
	contextRepo domain.ContextRepository
}

func NewGuestUsecase(gr domain.GuestRepository, cr domain.ContextRepository) GuestUsecase {
	return &guestUsecase{guestRepo: gr, contextRepo: cr}
}

func (u *guestUsecase) GetAllByContextID(ctx context.Context, contextID int64) ([]domain.Guest, error) {
	// Verify context exists
	c, err := u.contextRepo.GetByID(ctx, contextID)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, errors.New("context not found")
	}

	return u.guestRepo.GetAllByContextID(ctx, contextID)
}

func (u *guestUsecase) GetByID(ctx context.Context, contextID int64, guestID int64) (*domain.Guest, error) {
	return u.guestRepo.GetByID(ctx, contextID, guestID)
}

func (u *guestUsecase) Create(ctx context.Context, g *domain.Guest) error {
	c, err := u.contextRepo.GetByID(ctx, g.ContextID)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("context not found")
	}

	if g.Slug == "" {
		return errors.New("slug is required")
	}

	existing, err := u.guestRepo.GetBySlug(ctx, g.ContextID, g.Slug)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("guest with this slug already exists in this context")
	}

	return u.guestRepo.Create(ctx, g)
}

func (u *guestUsecase) Update(ctx context.Context, contextID int64, guestID int64, name, slug string) error {
	existing, err := u.guestRepo.GetByID(ctx, contextID, guestID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("guest not found")
	}

	if slug == "" {
		return errors.New("slug is required")
	}

	// If slug changed, verify uniqueness
	if existing.Slug != slug {
		dup, err := u.guestRepo.GetBySlug(ctx, contextID, slug)
		if err != nil {
			return err
		}
		if dup != nil {
			return errors.New("guest with this slug already exists in this context")
		}
	}

	existing.Name = name
	existing.Slug = slug

	return u.guestRepo.Update(ctx, existing)
}

func (u *guestUsecase) Delete(ctx context.Context, contextID int64, guestID int64) error {
	existing, err := u.guestRepo.GetByID(ctx, contextID, guestID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("guest not found")
	}
	return u.guestRepo.Delete(ctx, contextID, guestID)
}
