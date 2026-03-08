package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/util"
)

type Question struct {
	Repo ports.QuestionRepository
}

func (u *Question) Create(ctx context.Context, materialID uuid.UUID, data domain.CreateQuestionData) (uuid.UUID, error) {
	return u.Repo.Create(ctx, materialID, data)
}

func (u *Question) GetByMaterial(ctx context.Context, materialID uuid.UUID) ([]domain.Question, error) {
	return u.Repo.GetByMaterial(ctx, materialID)
}

func (u *Question) Move(ctx context.Context, id uuid.UUID, prevID *uuid.UUID, nextID *uuid.UUID) error {
	position, err := util.Midpoint(ctx, id, prevID, nextID, u.Repo.GetPosition)
	if err != nil {
		return err
	}

	return u.Repo.Update(ctx, id, domain.UpdateQuestionData{
		Position: &position,
	})
}

func (u *Question) Update(ctx context.Context, id uuid.UUID, data domain.UpdateQuestionData) error {
	return u.Repo.Update(ctx, id, data)
}
