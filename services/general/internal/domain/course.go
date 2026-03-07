package domain

import (
	"github.com/google/uuid"
)

type CourseLevel string

const (
	Beginner     CourseLevel = "beginner"
	Intermediate CourseLevel = "intermediate"
	Advanced     CourseLevel = "advanced"
)

type Course struct {
	ID                 uuid.UUID
	InstructorID       uuid.UUID
	Title              string
	Subtitle           *string
	Description        string
	Categories         []string
	Level              CourseLevel
	ThumbnailURL       *string
	PromoVideoURL      *string
	Price              float32
	Prerequisites      []uuid.UUID
	Requirements       []string
	LearningObjectives []string
	TargetAudiences    []string
}

type CreateCourseData struct {
	Title              string
	Description        string
	Categories         []string
	Level              CourseLevel
	Price              float32
	Prerequisites      []uuid.UUID
	LearningObjectives []string
}

type UpdateCourseData struct {
	Subtitle        *string
	ThumbnailURL    *string
	PromoVideoURL   *string
	Requirements    *[]string
	TargetAudiences *[]string
}
