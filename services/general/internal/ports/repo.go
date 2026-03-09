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

type MaterialRepository interface {
	Create(ctx context.Context, courseID uuid.UUID, data domain.CreateMaterialData) (uuid.UUID, error)
	GetByCourse(ctx context.Context, courseID uuid.UUID) ([]domain.Material, error)
	GetPosition(ctx context.Context, id uuid.UUID) (float64, error)
	Update(ctx context.Context, id uuid.UUID, data domain.UpdateMaterialData) error
}

type QuestionRepository interface {
	Create(ctx context.Context, materialID uuid.UUID, data domain.CreateQuestionData) (uuid.UUID, error)
	GetByMaterial(ctx context.Context, materialID uuid.UUID) ([]domain.Question, error)
	GetPosition(ctx context.Context, id uuid.UUID) (float64, error)
	Update(ctx context.Context, id uuid.UUID, data domain.UpdateQuestionData) error
}

type AttemptRepository interface {
	Create(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) (uuid.UUID, error)
	GetByMaterial(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) ([]domain.Attempt, error)
}
