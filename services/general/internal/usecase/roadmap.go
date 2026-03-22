package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/util"
)

type Roadmap struct {
	Repo ports.RoadmapRepository
}

func (u *Roadmap) Create(ctx context.Context, authorID uuid.UUID, data domain.CreateRoadmapData) (uuid.UUID, error) {
	return u.Repo.Create(ctx, authorID, data)
}

func (u *Roadmap) AddCourse(ctx context.Context, id uuid.UUID, courseID uuid.UUID) error {
	return u.Repo.AddCourse(ctx, id, courseID)
}

func (u *Roadmap) Apply(ctx context.Context, id uuid.UUID, studentID uuid.UUID) error {
	return u.Repo.Apply(ctx, id, studentID)
}

func (u *Roadmap) GetAll(ctx context.Context) ([]uuid.UUID, error) {
	return u.Repo.GetAll(ctx)
}

func (u *Roadmap) GetByStudent(ctx context.Context, studentID uuid.UUID) ([]uuid.UUID, error) {
	return u.Repo.GetByStudent(ctx, studentID)
}

func (u *Roadmap) Get(ctx context.Context, id uuid.UUID) (*domain.Roadmap, error) {
	return u.Repo.Get(ctx, id)
}

func (u *Roadmap) Move(ctx context.Context, id uuid.UUID, prevID *uuid.UUID, nextID *uuid.UUID) error {
	getPosition := func(ctx context.Context, id uuid.UUID) (float64, error) {
		roadmap, err := u.Repo.Get(ctx, id)
		if err != nil {
			return 0, err
		}

		return roadmap.Position, nil
	}
	position, err := util.Midpoint(ctx, id, prevID, nextID, getPosition)
	if err != nil {
		return err
	}

	return u.Repo.Update(ctx, id, domain.UpdateRoadmapData{Position: &position})
}

func (u *Roadmap) Update(ctx context.Context, id uuid.UUID, data domain.UpdateRoadmapData) error {
	return u.Repo.Update(ctx, id, data)
}
