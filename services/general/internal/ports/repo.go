package ports

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
)

type Repository interface {
	GetCourse(ctx context.Context, id uuid.UUID) (*domain.Course, error)
}
