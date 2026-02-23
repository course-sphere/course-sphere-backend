package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

var (
	ErrInvalidPrice = fmt.Errorf("price must be non negative")
	ErrNotOwner     = fmt.Errorf("User not own the resource")
)

type Course struct {
	CourseRepo   ports.CourseRepository
	MaterialRepo ports.MaterialRepository
}

func (u *Course) Create(ctx context.Context, instructorID uuid.UUID, course domain.CreateCourse) (uuid.UUID, error) {
	if course.Price < 0 {
		return uuid.Nil, ErrInvalidPrice
	}

	return u.CourseRepo.Create(ctx, instructorID, course)
}

func (u *Course) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	return u.CourseRepo.Get(ctx, id)
}

func (u *Course) AssertOwner(ctx context.Context, id uuid.UUID, instructorID uuid.UUID) error {
	isOwner, err := u.CourseRepo.IsOwner(ctx, id, instructorID)
	if err != nil {
		return err
	}
	if !isOwner {
		return ErrNotOwner
	}

	return nil
}

func (u *Course) GetProgress(ctx context.Context, id uuid.UUID, studentID uuid.UUID) (*domain.CourseProgress, error) {
	return u.CourseRepo.GetProgress(ctx, id, studentID)
}

func (u *Course) Update(
	ctx context.Context,
	id uuid.UUID,
	instructorID uuid.UUID,
	course domain.UpdateCourse,
) error {
	if err := u.AssertOwner(ctx, id, instructorID); err != nil {
		return err
	}

	if course.Price.TakeOr(0) < 0 {
		return ErrInvalidPrice
	}

	return u.CourseRepo.Update(ctx, id, course)
}

func (u *Course) Delete(ctx context.Context, id uuid.UUID, instructorID uuid.UUID) error {
	if err := u.AssertOwner(ctx, id, instructorID); err != nil {
		return err
	}

	return u.CourseRepo.Delete(ctx, id)
}

func (u *Course) CreateMaterial(ctx context.Context, id uuid.UUID, instructorID uuid.UUID, material domain.CreateMaterial) (uuid.UUID, error) {
	if err := u.AssertOwner(ctx, id, instructorID); err != nil {
		return uuid.Nil, err
	}

	return u.MaterialRepo.Create(ctx, id, material)
}

func (u *Course) GetMaterials(ctx context.Context, id uuid.UUID) ([]domain.Material, error) {
	return u.MaterialRepo.GetManyByCourse(ctx, id)
}

func (u *Course) AssertMaterialOwner(ctx context.Context, materialID uuid.UUID, instructorID uuid.UUID) error {
	isOwner, err := u.MaterialRepo.IsOwner(ctx, materialID, instructorID)
	if err != nil {
		return err
	}
	if !isOwner {
		return ErrNotOwner
	}

	return nil
}

func (u *Course) UpdateMaterial(
	ctx context.Context,
	materialID uuid.UUID,
	instructorID uuid.UUID,
	material domain.UpdateMaterial,
) error {
	if err := u.AssertMaterialOwner(ctx, materialID, instructorID); err != nil {
		return err
	}

	return u.MaterialRepo.Update(ctx, materialID, material)
}

func (u *Course) AddDependencies(ctx context.Context, materialID uuid.UUID, instructorID uuid.UUID, dependencies []uuid.UUID) error {
	if err := u.AssertMaterialOwner(ctx, materialID, instructorID); err != nil {
		return err
	}

	return u.MaterialRepo.AddDependencies(ctx, materialID, dependencies)
}

func (u *Course) DeleteMaterial(
	ctx context.Context,
	materialID uuid.UUID,
	instructorID uuid.UUID,
) error {
	if err := u.AssertMaterialOwner(ctx, materialID, instructorID); err != nil {
		return err
	}

	return u.MaterialRepo.Delete(ctx, materialID)
}
