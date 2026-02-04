package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

type Course struct {
	repo ports.Repository
}

func NewCourse(r ports.Repository) *Course {
	return &Course{repo: r}
}

func (u *Course) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	course, err := u.repo.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	return course, nil
}
