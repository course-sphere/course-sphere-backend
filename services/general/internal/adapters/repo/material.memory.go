package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

type MemoryMaterial struct{}

var _ ports.MaterialRepository = &MemoryMaterial{}

func NewMemoryMaterial() *MemoryMaterial {
	return &MemoryMaterial{}
}

func (r *MemoryMaterial) Create(ctx context.Context, courseID uuid.UUID, material domain.CreateMaterial) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (r *MemoryMaterial) GetManyByCourse(ctx context.Context, courseID uuid.UUID) ([]domain.Material, error) {
	return nil, nil
}

func (r *MemoryMaterial) IsOwner(ctx context.Context, id uuid.UUID, instructorID uuid.UUID) (bool, error) {
	return false, nil
}

func (r *MemoryMaterial) Update(ctx context.Context, id uuid.UUID, course domain.UpdateMaterial) error {
	return nil
}

func (r *MemoryMaterial) AddDependencies(ctx context.Context, id uuid.UUID, dependencies []uuid.UUID) error {
	return nil
}

func (r *MemoryMaterial) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
