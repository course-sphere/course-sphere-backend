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
)

type Course struct {
	CourseRepo ports.CourseRepository
}

func (u *Course) Create(ctx context.Context, instructorID uuid.UUID, data domain.CreateCourseData) (uuid.UUID, error) {
	if data.Price < 0 {
		return uuid.Nil, ErrInvalidPrice
	}

	return u.CourseRepo.Create(ctx, instructorID, data)
}

func (u *Course) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	return u.CourseRepo.Get(ctx, id)
}

func (u *Course) GetAll(ctx context.Context) ([]domain.Course, error) {
	return u.CourseRepo.GetAll(ctx)
}

func (u *Course) Update(ctx context.Context, id uuid.UUID, instructorID uuid.UUID, data domain.UpdateCourseData) error {
	return u.CourseRepo.Update(ctx, id, instructorID, data)
}
