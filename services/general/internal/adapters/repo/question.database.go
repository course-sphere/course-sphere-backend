package repo

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/adapters/repo/database"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jinzhu/copier"
)

type QuestionDatabase struct {
	Pool *pgxpool.Pool
}

var _ ports.QuestionRepository = &QuestionDatabase{}

func (db *QuestionDatabase) Create(ctx context.Context, materialID uuid.UUID, data domain.CreateQuestionData) (uuid.UUID, error) {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	inner := database.New(db.Pool).WithTx(tx)

	id, err := inner.CreateQuestion(ctx, database.CreateQuestionParams{
		MaterialID: materialID,
		Question:   data.Question,
	})
	if err != nil {
		return uuid.Nil, err
	}

	for _, answer := range data.PossibleAnswers {
		_, err := inner.CreateQuestionPossibleAnswer(ctx, database.CreateQuestionPossibleAnswerParams{
			QuestionID: id,
			Answer:     answer,
		})
		if err != nil {
			return uuid.Nil, err
		}
	}

	for _, criterion := range data.Criteria {
		_, err := inner.CreateQuestionCriterion(ctx, database.CreateQuestionCriterionParams{
			QuestionID: id,
			Criterion:  criterion.Criterion,
			Score:      criterion.Score,
		})
		if err != nil {
			return uuid.Nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (db *QuestionDatabase) GetByMaterial(ctx context.Context, materialID uuid.UUID) ([]domain.Question, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetQuestionsByMaterial(ctx, materialID)
	if err != nil {
		return nil, err
	}
	var questions []domain.Question
	copier.Copy(&questions, &raw)

	for i := range questions {
		questions[i].PossibleAnswers, err = inner.GetQuestionPossibleAnswers(ctx, questions[i].ID)
		if err != nil {
			return nil, err
		}

		criteria, err := inner.GetQuestionCriteria(ctx, questions[i].ID)
		if err != nil {
			return nil, err
		}
		copier.Copy(&questions[i].Criteria, &criteria)
	}

	return questions, nil
}

func (db *QuestionDatabase) Update(ctx context.Context, id uuid.UUID, data domain.UpdateQuestionData) error {
	inner := database.New(db.Pool)

	params := database.UpdateQuestionParams{ID: id}
	copier.Copy(&params, &data)

	return inner.UpdateQuestion(ctx, params)
}
