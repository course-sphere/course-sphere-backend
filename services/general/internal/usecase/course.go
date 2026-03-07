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

func (u *Course) Create(ctx context.Context, instructorID uuid.UUID, course domain.CreateCourseData) (uuid.UUID, error) {
	if course.Price < 0 {
		return uuid.Nil, ErrInvalidPrice
	}

	return u.CourseRepo.Create(ctx, instructorID, course)
}

func (u *Course) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	return u.CourseRepo.Get(ctx, id)
}
