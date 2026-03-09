package http

import (
	"github.com/google/uuid"

	sharedDomain "github.com/course-sphere/course-sphere-backend/shared/domain"
)

type CourseLevel string

const (
	Beginner     CourseLevel = "beginner"
	Intermediate CourseLevel = "intermediate"
	Advanced     CourseLevel = "advanced"
)

type CourseStatus string

const (
	Draft      CourseStatus = "draft"
	AIApproved CourseStatus = "ai-approved"
	Approved   CourseStatus = "approved"
	Removed    CourseStatus = "removed"
)

type Course struct {
	ID                 uuid.UUID         `json:"id" format:"uuid" description:"Unique identifier of the course" example:"550e8400-e29b-41d4-a716-446655440000"`
	Instructor         sharedDomain.User `json:"instructor" description:"Instructor user object"`
	Title              string            `json:"title" description:"Course title" example:"Introduction to Calculus"`
	Subtitle           *string           `json:"subtitle,omitempty" description:"Optional subtitle for the course" example:"Limits, derivatives and applications"`
	Description        string            `json:"description" description:"Detailed description of the course"`
	Categories         []string          `json:"categories,omitempty" description:"Course categories"`
	Level              CourseLevel       `json:"level" description:"Difficulty level of the course" enums:"beginner,intermediate,advanced" example:"beginner"`
	ThumbnailURL       *string           `json:"thumbnail_url,omitempty" format:"uri" description:"URL to the course thumbnail image"`
	PromoVideoURL      *string           `json:"promo_video_url,omitempty" format:"uri" description:"URL to the promotional video"`
	Price              float32           `json:"price" description:"Course price in USD" example:"49.99"`
	Prerequisites      []uuid.UUID       `json:"prerequisites,omitempty" format:"uuid" description:"IDs of prerequisite courses"`
	Requirements       []string          `json:"requirements,omitempty" description:"Requirements students should meet before taking the course"`
	LearningObjectives []string          `json:"learning_objectives,omitempty" description:"Skills or knowledge students will gain"`
	TargetAudiences    []string          `json:"target_audiences,omitempty" description:"Intended audience for the course"`
	Status             CourseStatus      `json:"status" description:"Status of the course" enums:"draft,ai-approved,approved,removed" example:"draft"`

	Total         int32 `json:"total" description:"Total number of learners enrolled" example:"120"`
	TotalRequired int32 `json:"total_required" description:"Total number of required materials/steps" example:"10"`
}

type CourseProgress struct {
	CompletedMaterials []uuid.UUID `json:"completed_materials,omitempty" format:"uuid" description:"IDs of completed materials (UUIDs)"`
	IsCompleted        bool        `json:"is_completed" description:"Whether the course is completed by the user" example:"false"`
}

type CreateCourseData struct {
	Title              string      `json:"title" description:"Course title" example:"Introduction to Calculus"`
	Description        string      `json:"description" description:"Detailed description of the course"`
	Categories         []string    `json:"categories,omitempty" description:"Course categories"`
	Level              CourseLevel `json:"level" description:"Difficulty level of the course" enums:"beginner,intermediate,advanced" example:"beginner"`
	Price              float32     `json:"price" description:"Course price in USD" example:"49.99"`
	Prerequisites      []uuid.UUID `json:"prerequisites,omitempty" format:"uuid" description:"IDs of prerequisite courses"`
	LearningObjectives []string    `json:"learning_objectives,omitempty" description:"Skills students will gain"`
}

type UpdateCourseData struct {
	Title              *string      `json:"title,omitempty" description:"Updated course title"`
	Subtitle           *string      `json:"subtitle,omitempty" description:"Updated course subtitle"`
	Description        *string      `json:"description,omitempty" description:"Updated course description"`
	Categories         *[]string    `json:"categories,omitempty" description:"Updated course categories"`
	Level              *CourseLevel `json:"level,omitempty" enums:"beginner,intermediate,advanced" description:"Updated course level"`
	ThumbnailURL       *string      `json:"thumbnail_url,omitempty" format:"uri" description:"Updated thumbnail URL"`
	PromoVideoURL      *string      `json:"promo_video_url,omitempty" format:"uri" description:"Updated promotional video URL"`
	Price              *float32     `json:"price,omitempty" description:"Updated course price"`
	Prerequisites      *[]uuid.UUID `json:"prerequisites,omitempty" format:"uuid" description:"Updated prerequisite courses"`
	Requirements       *[]string    `json:"requirements,omitempty" description:"Updated requirements"`
	LearningObjectives *[]string    `json:"learning_objectives,omitempty" description:"Updated learning objectives"`
	TargetAudiences    *[]string    `json:"target_audiences,omitempty" description:"Updated target audiences"`
	Status             CourseStatus `json:"status,omitempty" description:"Status of the course" enums:"draft,ai-approved,approved,removed" example:"draft"`
}
