package ports

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
)

type CourseRepository interface {
	Create(ctx context.Context, instructorID uuid.UUID, course domain.CreateCourseData) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.Course, error)
	IsOwner(ctx context.Context, id uuid.UUID, instructorID uuid.UUID) (bool, error)
	Update(ctx context.Context, id uuid.UUID, course domain.UpdateCourseData) error
	Delete(ctx context.Context, id uuid.UUID) error
}
