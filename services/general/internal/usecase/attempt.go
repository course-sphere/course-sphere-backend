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

func (u *Attempt) GetByMaterial(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) ([]domain.Attempt, error) {
	return u.Repo.GetByMaterial(ctx, materialID, studentID)
}
