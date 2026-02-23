package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/ports"
)

type MemoryCourse struct {
	courses map[uuid.UUID]domain.Course
}

var _ ports.CourseRepository = &MemoryCourse{}

func NewMemoryCourse() *MemoryCourse {
	return &MemoryCourse{
		courses: make(map[uuid.UUID]domain.Course),
	}
}

func (r *MemoryCourse) Create(ctx context.Context, instructorID uuid.UUID, data domain.CreateCourse) (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	course := domain.Course{
		ID:           id,
		InstructorID: instructorID,
	}
	copier.Copy(&course, &data)

	r.courses[id] = course

	return id, nil
}

func (r *MemoryCourse) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	course, exists := r.courses[id]
	if !exists {
		return nil, ErrNotFound
	}

	return &course, nil
}

func (r *MemoryCourse) IsOwner(ctx context.Context, id uuid.UUID, instructorID uuid.UUID) (bool, error) {
	course, exists := r.courses[id]
	if !exists {
		return false, ErrNotFound
	}

	return course.InstructorID == instructorID, nil
}

func (r *MemoryCourse) GetProgress(ctx context.Context, id uuid.UUID, studentID uuid.UUID) (*domain.CourseProgress, error) {
	return &domain.CourseProgress{Progress: .92, IsPassed: true}, nil
}

func (r *MemoryCourse) Update(ctx context.Context, id uuid.UUID, data domain.UpdateCourse) error {
	course, exists := r.courses[id]
	if !exists {
		return ErrNotFound
	}

	err := copier.CopyWithOption(&course, &data, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return err
	}
	r.courses[id] = course

	return nil
}

func (r *MemoryCourse) Delete(ctx context.Context, id uuid.UUID) error {
	delete(r.courses, id)

	return nil
}
