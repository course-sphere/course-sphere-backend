package usecase

import (
	"context"
	"slices"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

type Attempt struct {
	Repo         ports.AttemptRepository
	MaterialRepo ports.MaterialRepository
	QuestionRepo ports.QuestionRepository
}

func (u *Attempt) Create(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) (uuid.UUID, error) {
	return u.Repo.Create(ctx, materialID, studentID)
}

func (u *Attempt) CreateDetails(ctx context.Context, id uuid.UUID, data []domain.CreateAttemptDetailData) error {
	err := u.Repo.CreateDetails(ctx, id, data)
	if err != nil {
		return err
	}

	attempt, err := u.Repo.Get(ctx, id)
	if err != nil {
		return err
	}
	material, err := u.MaterialRepo.Get(ctx, attempt.MaterialID)
	if err != nil {
		return err
	}
	if material.Kind != domain.QuizMaterial {
		return nil
	}
	questions, err := u.QuestionRepo.GetByMaterial(ctx, material.ID)
	if err != nil {
		return err
	}

	score := int32(0)

	for _, item := range data {
		i := slices.IndexFunc(questions, func(question domain.Question) bool {
			return question.ID == item.QuestionID
		})
		for _, criterion := range questions[i].Criteria {
			if item.Answer == criterion.Criterion {
				score += criterion.Score
			}
		}
	}

	return u.Repo.Update(ctx, id, score)
}

func (u *Attempt) GetDetails(ctx context.Context, id uuid.UUID) ([]domain.AttemptDetail, error) {
	return u.Repo.GetDetails(ctx, id)
}

func (u *Attempt) GetByMaterial(ctx context.Context, materialID uuid.UUID, studentID uuid.UUID) ([]domain.Attempt, error) {
	return u.Repo.GetByMaterial(ctx, materialID, studentID)
}

func (u *Attempt) Update(ctx context.Context, id uuid.UUID, score int32) error {
	return u.Repo.Update(ctx, id, score)
}
