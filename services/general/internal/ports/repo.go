package ports

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
)

type CourseRepository interface {
	Create(ctx context.Context, instructorID uuid.UUID, data domain.CreateCourseData) (uuid.UUID, error)
	Enroll(ctx context.Context, id uuid.UUID, studentID uuid.UUID) error
	Get(ctx context.Context, id uuid.UUID) (*domain.Course, error)
	GetAll(ctx context.Context) ([]domain.Course, error)
	GetEnrolled(ctx context.Context, studentID uuid.UUID) ([]uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, instructorID uuid.UUID, data domain.UpdateCourseData) error
}

type MaterialRepository interface {
	Create(ctx context.Context, courseID uuid.UUID, data domain.CreateMaterialData) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.Material, error)
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
	CreateDetails(ctx context.Context, id uuid.UUID, data []domain.CreateAttemptDetailData) error
	GetByMaterial(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) ([]domain.Attempt, error)
	GetDetails(ctx context.Context, id uuid.UUID) ([]domain.AttemptDetail, error)
	Update(ctx context.Context, id uuid.UUID, score int32) error
}

type RoadmapRepository interface {
	Create(ctx context.Context, authorID uuid.UUID, data domain.CreateRoadmapData) (uuid.UUID, error)
	AddCourse(ctx context.Context, id uuid.UUID, courseID uuid.UUID) error
	Apply(ctx context.Context, id uuid.UUID, studentID uuid.UUID) error
	GetAll(ctx context.Context) ([]uuid.UUID, error)
	GetByStudent(ctx context.Context, studentID uuid.UUID) ([]uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (*domain.Roadmap, error)
	Update(ctx context.Context, id uuid.UUID, data domain.UpdateRoadmapData) error
}
