package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jinzhu/copier"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/adapters/repo/database"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

type RoadmapDatabase struct {
	Pool *pgxpool.Pool
}

var _ ports.RoadmapRepository = &RoadmapDatabase{}

func (db *RoadmapDatabase) Create(ctx context.Context, authorID uuid.UUID, data domain.CreateRoadmapData) (uuid.UUID, error) {
	inner := database.New(db.Pool)

	return inner.CreateRoadmap(ctx, database.CreateRoadmapParams{
		AuthorID:    authorID,
		Title:       data.Title,
		Description: data.Description,
	})
}

func (db *RoadmapDatabase) AddCourse(ctx context.Context, id uuid.UUID, courseID uuid.UUID) error {
	inner := database.New(db.Pool)

	return inner.AddRoadmapCourse(ctx, database.AddRoadmapCourseParams{
		ID:       id,
		CourseID: courseID,
	})
}

func (db *RoadmapDatabase) Apply(ctx context.Context, id uuid.UUID, studentID uuid.UUID) error {
	inner := database.New(db.Pool)

	return inner.ApplyRoadmap(ctx, database.ApplyRoadmapParams{
		ID:        id,
		StudentID: studentID,
	})
}

func (db *RoadmapDatabase) GetAll(ctx context.Context) ([]uuid.UUID, error) {
	inner := database.New(db.Pool)

	return inner.GetAllRoadmaps(ctx)
}

func (db *RoadmapDatabase) GetByStudent(ctx context.Context, studentID uuid.UUID) ([]uuid.UUID, error) {
	inner := database.New(db.Pool)

	return inner.GetRoadmapsByStudent(ctx, studentID)
}

func (db *RoadmapDatabase) Get(ctx context.Context, id uuid.UUID) (*domain.Roadmap, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetRoadmap(ctx, id)
	if err != nil {
		return nil, err
	}

	var roadmap domain.Roadmap
	copier.Copy(&roadmap, &raw)

	courses, err := inner.GetRoadmapCourse(ctx, id)
	copier.Copy(&roadmap.Courses, &courses)

	return &roadmap, nil
}

func (db *RoadmapDatabase) Update(ctx context.Context, id uuid.UUID, data domain.UpdateRoadmapData) error {
	inner := database.New(db.Pool)

	return inner.UpdateRoadmap(ctx, database.UpdateRoadmapParams{
		ID:          id,
		Position:    data.Position,
		Title:       data.Title,
		Description: data.Description,
	})
}
