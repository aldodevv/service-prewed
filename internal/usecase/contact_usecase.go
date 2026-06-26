package usecase

import (
	"context"
	"errors"
	"service-wedding/internal/domain"
)

type ContactUsecase interface {
	Create(ctx context.Context, msg *domain.ContactMessage) error
	GetAll(ctx context.Context) ([]domain.ContactMessage, error)
}

type contactUsecase struct {
	contactRepo domain.ContactRepository
}

func NewContactUsecase(cr domain.ContactRepository) ContactUsecase {
	return &contactUsecase{contactRepo: cr}
}

func (u *contactUsecase) Create(ctx context.Context, msg *domain.ContactMessage) error {
	if msg.Name == "" || msg.Email == "" || msg.Message == "" {
		return errors.New("name, email, and message are required")
	}
	return u.contactRepo.Create(ctx, msg)
}

func (u *contactUsecase) GetAll(ctx context.Context) ([]domain.ContactMessage, error) {
	return u.contactRepo.GetAll(ctx)
}
