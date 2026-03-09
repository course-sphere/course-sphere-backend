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

func (db *AttemptDatabase) Create(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) (uuid.UUID, error) {
	inner := database.New(db.Pool)

	return inner.CreateAttempt(ctx, database.CreateAttemptParams{
		MaterialID: materialID,
		StudentID:  studentID,
	})
}

func (db *AttemptDatabase) CreateDetails(ctx context.Context, id uuid.UUID, data []domain.CreateAttemptDetailData) error {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	inner := database.New(db.Pool).WithTx(tx)

	for _, attempt := range data {
		err := inner.CreateAttemptDetail(ctx, database.CreateAttemptDetailParams{
			ID:         id,
			QuestionID: attempt.QuestionID,
			Answer:     attempt.Answer,
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
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

func (db *AttemptDatabase) GetDetails(ctx context.Context, id uuid.UUID) ([]domain.AttemptDetail, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetAttemptDetails(ctx, id)
	if err != nil {
		return nil, err
	}
	var details []domain.AttemptDetail
	copier.Copy(&details, &raw)

	return details, nil
}

func (db *AttemptDatabase) Update(ctx context.Context, id uuid.UUID, score int32) error {
	inner := database.New(db.Pool)

	return inner.UpdateAttempt(ctx, database.UpdateAttemptParams{
		ID:    id,
		Score: &score,
	})
}
