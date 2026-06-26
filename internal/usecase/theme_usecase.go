package usecase

import (
	"context"
	"errors"
	"service-wedding/internal/domain"
)

type ThemeUsecase interface {
	GetAll(ctx context.Context) ([]domain.Theme, error)
	GetByID(ctx context.Context, id int64) (*domain.Theme, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Theme, error)
	Create(ctx context.Context, theme *domain.Theme) error
	Update(ctx context.Context, id int64, theme *domain.Theme) error
	Delete(ctx context.Context, id int64) error
}

type themeUsecase struct {
	themeRepo domain.ThemeRepository
}

func NewThemeUsecase(tr domain.ThemeRepository) ThemeUsecase {
	return &themeUsecase{themeRepo: tr}
}

func (u *themeUsecase) GetAll(ctx context.Context) ([]domain.Theme, error) {
	return u.themeRepo.GetAll(ctx)
}

func (u *themeUsecase) GetByID(ctx context.Context, id int64) (*domain.Theme, error) {
	return u.themeRepo.GetByID(ctx, id)
}

func (u *themeUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Theme, error) {
	return u.themeRepo.GetBySlug(ctx, slug)
}

func (u *themeUsecase) Create(ctx context.Context, t *domain.Theme) error {
	if t.Slug == "" {
		return errors.New("slug is required")
	}
	existing, err := u.themeRepo.GetBySlug(ctx, t.Slug)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("theme with this slug already exists")
	}
	return u.themeRepo.Create(ctx, t)
}

func (u *themeUsecase) Update(ctx context.Context, id int64, t *domain.Theme) error {
	existing, err := u.themeRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("theme not found")
	}

	t.ID = id
	return u.themeRepo.Update(ctx, t)
}

func (u *themeUsecase) Delete(ctx context.Context, id int64) error {
	existing, err := u.themeRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("theme not found")
	}
	return u.themeRepo.Delete(ctx, id)
}
