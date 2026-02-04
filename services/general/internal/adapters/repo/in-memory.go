package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

type InMemory struct{}

func NewInMemory() ports.Repository {
	return &InMemory{}
}

func (r *InMemory) GetCourse(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
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

var _ ports.Repository = &InMemory{}
var _ ports.Repository = (*InMemory)(nil)
