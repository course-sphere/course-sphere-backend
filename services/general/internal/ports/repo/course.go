package repo

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/google/uuid"
)

type CourseRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Course, error)
}
