package ports

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
)

type CourseRepository interface {
	Create(ctx context.Context, instructorID uuid.UUID, course domain.CreateCourse) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.Course, error)
	IsOwner(ctx context.Context, id uuid.UUID, instructorID uuid.UUID) (bool, error)
	GetProgress(ctx context.Context, id uuid.UUID, studentID uuid.UUID) (*domain.CourseProgress, error)
	Update(ctx context.Context, id uuid.UUID, course domain.UpdateCourse) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type MaterialRepository interface {
	Create(ctx context.Context, courseID uuid.UUID, material domain.CreateMaterial) (uuid.UUID, error)
	GetManyByCourse(ctx context.Context, courseID uuid.UUID) ([]domain.Material, error)
	IsOwner(ctx context.Context, id uuid.UUID, instructorID uuid.UUID) (bool, error)
	Update(ctx context.Context, id uuid.UUID, course domain.UpdateMaterial) error
	AddDependencies(ctx context.Context, id uuid.UUID, dependencies []uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}
