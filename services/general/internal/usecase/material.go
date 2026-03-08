package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/util"
)

type Material struct {
	MaterialRepo ports.MaterialRepository
}

func (u *Material) Create(ctx context.Context, courseID uuid.UUID, data domain.CreateMaterialData) (uuid.UUID, error) {
	return u.MaterialRepo.Create(ctx, courseID, data)
}

func (u *Material) GetByCourse(ctx context.Context, courseID uuid.UUID) ([]domain.Material, error) {
	return u.MaterialRepo.GetByCourse(ctx, courseID)
}

func (u *Material) Move(ctx context.Context, id uuid.UUID, prevID *uuid.UUID, nextID *uuid.UUID) error {
	position, err := util.Midpoint(ctx, id, prevID, nextID, u.MaterialRepo.GetPosition)
	if err != nil {
		return err
	}

	return u.MaterialRepo.Update(ctx, id, domain.UpdateMaterialData{
		Position: &position,
	})
}

func (u *Material) Update(ctx context.Context, id uuid.UUID, data domain.UpdateMaterialData) error {
	return u.MaterialRepo.Update(ctx, id, data)
}
