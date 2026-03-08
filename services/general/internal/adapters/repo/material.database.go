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

type MaterialDatabase struct {
	Pool *pgxpool.Pool
}

var _ ports.MaterialRepository = &MaterialDatabase{}

func (db *MaterialDatabase) Create(ctx context.Context, courseID uuid.UUID, data domain.CreateMaterialData) (uuid.UUID, error) {
	inner := database.New(db.Pool)

	params := database.CreateMaterialParams{CourseID: courseID}
	copier.Copy(&params, &data)

	return inner.CreateMaterial(ctx, params)
}

func (db *MaterialDatabase) CreateAttempt(ctx context.Context, id uuid.UUID, studentID uuid.UUID, score *int32) (uuid.UUID, error) {
	inner := database.New(db.Pool)

	return inner.CreateAttempt(ctx, database.CreateAttemptParams{
		MaterialID: id,
		StudentID:  studentID,
		Score:      score,
	})
}

func (db *MaterialDatabase) GetByCourse(ctx context.Context, courseID uuid.UUID) ([]domain.Material, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetMaterialsByCourse(ctx, courseID)
	if err != nil {
		return nil, err
	}

	var materials []domain.Material
	copier.Copy(&materials, &raw)

	return materials, nil
}

func (db *MaterialDatabase) GetPosition(ctx context.Context, id uuid.UUID) (float64, error) {
	inner := database.New(db.Pool)

	return inner.GetMaterialPosition(ctx, id)
}

func (db *MaterialDatabase) GetAttempts(ctx context.Context, id uuid.UUID, studentID uuid.UUID) ([]domain.Attempt, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetAttemptsByMaterial(ctx, database.GetAttemptsByMaterialParams{
		MaterialID: id,
		StudentID:  studentID,
	})
	if err != nil {
		return nil, err
	}
	var attempts []domain.Attempt
	copier.Copy(&attempts, &raw)

	return attempts, nil
}

func (db *MaterialDatabase) Update(ctx context.Context, id uuid.UUID, data domain.UpdateMaterialData) error {
	inner := database.New(db.Pool)

	params := database.UpdateMaterialParams{ID: id}
	copier.Copy(&params, &data)

	return inner.UpdateMaterial(ctx, params)
}
