package domain

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
	Kind          MaterialKind
	Lesson        string
	Title         string
	Content       string
	RequiredScore *int64
	RequiredPeers *int64
}

type CreateMaterialData struct {
	Kind          MaterialKind
	Lesson        string
	Title         string
	Content       string
	RequiredScore *int64
	RequiredPeers *int64
}

type UpdateMaterialData struct {
	Lesson        *string
	Title         *string
	Content       *string
	RequiredScore *int64
	RequiredPeers *int64
}

type Question struct {
	Question        string
	PossibleAnswers []string
}

type QuestionCriterion struct {
	Criterion string
	Score     int64
}
