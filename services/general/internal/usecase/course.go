package usecase

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

var (
	ErrInvalidPrice = fmt.Errorf("price must be non negative")
)

type Course struct {
	Repo         ports.CourseRepository
	MaterialRepo ports.MaterialRepository
	AttemptRepo  ports.AttemptRepository
}

func (u *Course) Create(ctx context.Context, instructorID uuid.UUID, data domain.CreateCourseData) (uuid.UUID, error) {
	if data.Price < 0 {
		return uuid.Nil, ErrInvalidPrice
	}

	return u.Repo.Create(ctx, instructorID, data)
}

func (u *Course) Enroll(ctx context.Context, id uuid.UUID, studentID uuid.UUID) error {
	return u.Repo.Enroll(ctx, id, studentID)
}

func (u *Course) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	return u.Repo.Get(ctx, id)
}

func (u *Course) GetAll(ctx context.Context) ([]domain.Course, error) {
	return u.Repo.GetAll(ctx)
}

// TODO: optimize
func (u *Course) GetProgress(ctx context.Context, id uuid.UUID, studentID uuid.UUID) (*domain.CourseProgress, error) {
	course, err := u.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	materials, err := u.MaterialRepo.GetByCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	completed := make([]uuid.UUID, 0)
	requiredCompleted := int32(0)
	for _, material := range materials {
		attemps, err := u.AttemptRepo.GetByMaterial(ctx, material.ID, studentID)
		if err != nil {
			return nil, err
		}
		if len(attemps) == 0 {
			continue
		}

		if material.Kind == domain.QuizMaterial || material.Kind == domain.AssignmentMaterial {
			maxScore := *slices.MaxFunc(attemps, func(a, b domain.Attempt) int {
				return cmp.Compare(*a.Score, *b.Score)
			}).Score
			if maxScore < *material.RequiredScore {
				continue
			}
		}

		completed = append(completed, material.ID)
		if material.IsRequired {
			requiredCompleted += 1
		}
	}

	return &domain.CourseProgress{
		CompletedMaterials: completed,
		IsCompleted:        requiredCompleted >= course.TotalRequired,
	}, nil
}

func (u *Course) Update(ctx context.Context, id uuid.UUID, instructorID uuid.UUID, data domain.UpdateCourseData) error {
	return u.Repo.Update(ctx, id, instructorID, data)
}
