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

type AttemptDatabase struct {
	Pool *pgxpool.Pool
}

var _ ports.AttemptRepository = &AttemptDatabase{}

func (db *AttemptDatabase) Create(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID, score *int32) (uuid.UUID, error) {
	inner := database.New(db.Pool)

	return inner.CreateAttempt(ctx, database.CreateAttemptParams{
		MaterialID: materialID,
		StudentID:  studentID,
		Score:      score,
	})
}

func (db *AttemptDatabase) GetByMaterial(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) ([]domain.Attempt, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetAttemptsByMaterial(ctx, database.GetAttemptsByMaterialParams{
		MaterialID: materialID,
		StudentID:  studentID,
	})
	if err != nil {
		return nil, err
	}
	var attempts []domain.Attempt
	copier.Copy(&attempts, &raw)

	return attempts, nil
}
