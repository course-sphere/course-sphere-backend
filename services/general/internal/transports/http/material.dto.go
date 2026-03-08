package http

import "github.com/google/uuid"

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

type CreateMaterialData struct {
	Kind          MaterialKind
	Lesson        string
	Title         string
	Content       *string
	RequiredScore *int32
	RequiredPeers *int32
	IsRequired    bool
}

type MoveMaterialData struct {
	PrevID *uuid.UUID
	NextID *uuid.UUID
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
