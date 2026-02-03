package repo

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/google/uuid"
)

type CourseRepository struct {
}

func (r *CourseRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	return &domain.Course{
		Id:         id,
		Thumbnail:  "/placeholder.jpg",
		Title:      "Course",
		Tags:       make([]string, 0),
		Instructor: "Instructor",
		Rating:     5.,
		Reviews:    100,
		Students:   100,
		Price:      10000,
	}, nil
}
