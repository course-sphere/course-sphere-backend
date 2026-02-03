package usecase

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports/repo"
	"github.com/google/uuid"
)

type CourseUsecase struct {
	repo repo.CourseRepository
}

func NewCourseUsecase(r repo.CourseRepository) *CourseUsecase {
	return &CourseUsecase{repo: r}
}

func (u *CourseUsecase) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	course, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return course, nil
}
