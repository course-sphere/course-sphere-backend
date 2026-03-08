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
	Position      int64
	Kind          MaterialKind
	Lesson        string
	Title         string
	Content       *string
	RequiredScore *int64
	RequiredPeers *int64
	IsRequired    bool
	Question      []Question
}

type CreateMaterialData struct {
	CourseID      uuid.UUID
	Kind          MaterialKind
	Lesson        string
	Title         string
	Content       *string
	RequiredScore *int64
	RequiredPeers *int64
	IsRequired    bool
}

type UpdateMaterialData struct {
	Position      *int64
	Lesson        *string
	Title         *string
	Content       *string
	RequiredScore *int64
	RequiredPeers *int64
	IsRequired    *bool
}

type QuestionCriterion struct {
	Criterion string
	Score     int64
}

type Question struct {
	Question        string
	Position        int64
	PossibleAnswers []string
	Criteria        []QuestionCriterion
}

type CreateQuestionData struct {
	Question        string
	PossibleAnswers []string
	Criteria        []QuestionCriterion
}

type UpadteQuestionData struct {
	Question        *string
	Position        *int64
	PossibleAnswers *[]string
	Criteria        *[]QuestionCriterion
}
