package http

import "github.com/google/uuid"

type CourseLevel string

const (
	Beginner     CourseLevel = "beginner"
	Intermediate CourseLevel = "intermediate"
	Advanced     CourseLevel = "advanced"
)

type Course struct {
	ID                 uuid.UUID   `json:"id"`
	Instructor         string      `json:"instructor"`
	Title              string      `json:"title"`
	Subtitle           *string     `json:"subtitle,omitempty"`
	Description        string      `json:"description"`
	Categories         []string    `json:"categories,omitempty"`
	Level              CourseLevel `json:"level"`
	ThumbnailURL       *string     `json:"thumbnail_url,omitempty"`
	PromoVideoURL      *string     `json:"promo_video_url,omitempty"`
	Price              float32     `json:"price"`
	Prerequisites      []uuid.UUID `json:"prerequisites,omitempty"`
	Requirements       []string    `json:"requirements,omitempty"`
	LearningObjectives []string    `json:"learning_objectives,omitempty"`
	TargetAudiences    []string    `json:"target_audiences,omitempty"`
}

type CreateCourseData struct {
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	Categories         []string    `json:"categories,omitempty"`
	Level              CourseLevel `json:"level"`
	Price              float32     `json:"price"`
	Prerequisites      []uuid.UUID `json:"prerequisites,omitempty"`
	LearningObjectives []string    `json:"learning_objectives,omitempty"`
}
