package domain

import (
	"github.com/google/uuid"
	"github.com/moznion/go-optional"
)

type Course struct {
	ID           uuid.UUID
	Thumbnail    string
	Title        string
	Tags         []string
	InstructorID uuid.UUID
	Price        float32
}

type CreateCourse struct {
	Thumbnail string
	Title     string
	Tags      []string
	Price     float32
}

type UpdateCourse struct {
	Thumbnail optional.Option[string]
	Title     optional.Option[string]
	Tags      optional.Option[[]string]
	Price     optional.Option[float32]
}

type CourseProgress struct {
	Progress float32
	IsPassed bool
}

type MaterialKind string

const (
	TextMaterial       MaterialKind = "text"
	VideoMaterial      MaterialKind = "video"
	QuizMaterial       MaterialKind = "quiz"
	AssignmentMaterial MaterialKind = "assignment"
)

type Material struct {
	ID            uuid.UUID
	Kind          MaterialKind
	Group         string
	Title         string
	Content       string
	Dependencies  []uuid.UUID
	RequiredScore int64
	RequiredPeers int64
	IsOptional    bool
}

type CreateMaterial struct {
	Kind          MaterialKind
	Group         string
	Title         string
	Content       string
	RequiredScore int64
	RequiredPeers int64
	IsOptional    bool
}

type UpdateMaterial struct {
	Group         optional.Option[string]
	Title         optional.Option[string]
	Content       optional.Option[string]
	RequiredScore optional.Option[int64]
	RequiredPeers optional.Option[int64]
	IsOptional    optional.Option[bool]
}
