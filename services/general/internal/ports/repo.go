package ports

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
)

type CourseRepository interface {
	Create(ctx context.Context, instructorID uuid.UUID, data domain.CreateCourseData) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.Course, error)
	GetAll(ctx context.Context) ([]domain.Course, error)
	Update(ctx context.Context, id uuid.UUID, instructorID uuid.UUID, data domain.UpdateCourseData) error
}
