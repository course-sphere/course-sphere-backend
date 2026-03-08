package domain

import (
	"time"

	"github.com/google/uuid"
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
	ID            uuid.UUID
	Position      float64
	Kind          MaterialKind
	Lesson        string
	Title         string
	Content       *string
	RequiredScore *int32
	RequiredPeers *int32
	IsRequired    bool
}

type MaterialAttempt struct {
	ID         uuid.UUID
	MaterialID uuid.UUID
	StudentID  uuid.UUID
	Score      *int32
	CreatedAt  time.Time
}

type CreateMaterialData struct {
	Kind          MaterialKind
	Lesson        string
	Title         string
	Content       *string
	RequiredScore *int32
	RequiredPeers *int32
	IsRequired    bool
}

type UpdateMaterialData struct {
	Position      *float64
	Lesson        *string
	Title         *string
	Content       *string
	RequiredScore *int32
	RequiredPeers *int32
	IsRequired    *bool
}
