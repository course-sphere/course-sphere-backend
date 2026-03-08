package http

import (
	"time"

	"github.com/google/uuid"

	sharedDomain "github.com/course-sphere/course-sphere-backend/shared/domain"
)

type MaterialKind string

const (
	TextMaterial       MaterialKind = "text"
	FileMaterial       MaterialKind = "file"
	VideoMaterial      MaterialKind = "video"
	QuizMaterial       MaterialKind = "quiz"
	AssignmentMaterial MaterialKind = "assignment"
)

type Material struct {
	ID            uuid.UUID    `json:"id" format:"uuid" description:"Unique material identifier (UUID v4)" example:"550e8400-e29b-41d4-a716-446655440000"`
	Position      float64      `json:"position" description:"Position/order of the material within the lesson (can be fractional)" example:"1.0"`
	Kind          MaterialKind `json:"kind" description:"Material type" enums:"text,file,video,quiz,assignment" example:"text"`
	Lesson        uuid.UUID    `json:"lesson" format:"uuid" description:"ID of the lesson this material belongs to (UUID)" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title         string       `json:"title" description:"Material title" example:"Introduction to limits"`
	Content       *string      `json:"content,omitempty" description:"Text content of the material (if applicable)" example:"This is the lesson content..."`
	RequiredScore *int32       `json:"required_score,omitempty" description:"Minimum score required (if applicable)" example:"70"`
	RequiredPeers *int32       `json:"required_peers,omitempty" description:"Number of peer reviews required (if applicable)" example:"3"`
	IsRequired    bool         `json:"is_required" description:"Whether completing this material is required" example:"true"`
}

type Attempt struct {
	ID         uuid.UUID         `json:"id" format:"uuid" description:"Unique attempt identifier (UUID)" example:"550e8400-e29b-41d4-a716-446655440000"`
	MaterialID uuid.UUID         `json:"material_id" format:"uuid" description:"ID of the material attempted" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Student    sharedDomain.User `json:"student" description:"Student who made the attempt"`
	Score      *int32            `json:"score,omitempty" description:"Score achieved in the attempt" example:"85"`
	CreatedAt  time.Time         `json:"created_at" format:"date-time" description:"Attempt creation timestamp" example:"2026-03-08T22:00:00Z"`
}

type CreateMaterialData struct {
	Kind          MaterialKind `json:"kind" description:"Material type" enums:"text,file,video,quiz,assignment" example:"text"`
	Lesson        uuid.UUID    `json:"lesson" format:"uuid" description:"Lesson ID the material will belong to (UUID)" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title         string       `json:"title" description:"Material title" example:"Introduction to limits"`
	Content       *string      `json:"content,omitempty" description:"Text content of the material (if applicable)"`
	RequiredScore *int32       `json:"required_score,omitempty" description:"Minimum score required (if applicable)" example:"70"`
	RequiredPeers *int32       `json:"required_peers,omitempty" description:"Number of peer reviews required (if applicable)" example:"3"`
	IsRequired    bool         `json:"is_required" description:"Whether completing this material is required" example:"false"`
}

type CreateAttemptData struct {
	Score *int32 `json:"score,omitempty" description:"Score submitted for the attempt" example:"85"`
}

type MoveMaterialData struct {
	PrevID *uuid.UUID `json:"prev_id,omitempty" format:"uuid" description:"UUID of the previous material (nil to place at beginning)"`
	NextID *uuid.UUID `json:"next_id,omitempty" format:"uuid" description:"UUID of the next material (nil to place at end)"`
}

type UpdateMaterialData struct {
	Position      *float64   `json:"position,omitempty" description:"Updated position/order" example:"2.5"`
	Lesson        *uuid.UUID `json:"lesson,omitempty" format:"uuid" description:"Updated lesson ID (UUID)"`
	Title         *string    `json:"title,omitempty" description:"Updated title"`
	Content       *string    `json:"content,omitempty" description:"Updated content"`
	RequiredScore *int32     `json:"required_score,omitempty" description:"Updated minimum score required" example:"75"`
	RequiredPeers *int32     `json:"required_peers,omitempty" description:"Updated number of required peer reviews" example:"2"`
	IsRequired    *bool      `json:"is_required,omitempty" description:"Updated required flag"`
}
