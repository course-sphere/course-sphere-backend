package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

type Attempt struct {
	Repo ports.AttemptRepository
}

func (u *Attempt) Create(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) (uuid.UUID, error) {
	return u.Repo.Create(ctx, materialID, studentID)
}

func (u *Attempt) CreateDetails(ctx context.Context, id uuid.UUID, data []domain.CreateAttemptDetailData) error {
	err := u.Repo.CreateDetails(ctx, id, data)
	if err != nil {
		return err
	}

	// TODO: implement
	return u.Repo.Update(ctx, id, 100)
}

func (u *Attempt) GetDetails(ctx context.Context, id uuid.UUID) ([]domain.AttemptDetail, error) {
	return u.Repo.GetDetails(ctx, id)
}

func (u *Attempt) GetByMaterial(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) ([]domain.Attempt, error) {
	return u.Repo.GetByMaterial(ctx, materialID, studentID)
}
