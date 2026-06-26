package usecase

import (
	"context"
	"errors"
	"service-wedding/internal/domain"
)

type ContextUsecase interface {
	GetAll(ctx context.Context) ([]domain.Context, error)
	GetByID(ctx context.Context, id int64) (*domain.Context, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Context, error)
	Create(ctx context.Context, c *domain.Context) error
	Update(ctx context.Context, id int64, c *domain.Context) error
	Delete(ctx context.Context, id int64) error
}

type contextUsecase struct {
	contextRepo domain.ContextRepository
	themeRepo   domain.ThemeRepository
}

func NewContextUsecase(cr domain.ContextRepository, tr domain.ThemeRepository) ContextUsecase {
	return &contextUsecase{contextRepo: cr, themeRepo: tr}
}

func (u *contextUsecase) GetAll(ctx context.Context) ([]domain.Context, error) {
	return u.contextRepo.GetAll(ctx)
}

func (u *contextUsecase) GetByID(ctx context.Context, id int64) (*domain.Context, error) {
	return u.contextRepo.GetByID(ctx, id)
}

func (u *contextUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Context, error) {
	return u.contextRepo.GetBySlug(ctx, slug)
}

func (u *contextUsecase) Create(ctx context.Context, c *domain.Context) error {
	if c.Slug == "" {
		return errors.New("slug is required")
	}
	existing, err := u.contextRepo.GetBySlug(ctx, c.Slug)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("context with this slug already exists")
	}

	// Validate theme_id
	if c.ThemeID > 0 {
		theme, err := u.themeRepo.GetByID(ctx, c.ThemeID)
		if err != nil {
			return err
		}
		if theme == nil {
			return errors.New("selected theme not found")
		}
	}

	return u.contextRepo.Create(ctx, c)
}

func (u *contextUsecase) Update(ctx context.Context, id int64, c *domain.Context) error {
	existing, err := u.contextRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("context not found")
	}

	// Validate theme_id
	if c.ThemeID > 0 {
		theme, err := u.themeRepo.GetByID(ctx, c.ThemeID)
		if err != nil {
			return err
		}
		if theme == nil {
			return errors.New("selected theme not found")
		}
	}

	c.ID = id
	return u.contextRepo.Update(ctx, c)
}

func (u *contextUsecase) Delete(ctx context.Context, id int64) error {
	existing, err := u.contextRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("context not found")
	}
	return u.contextRepo.Delete(ctx, id)
}
