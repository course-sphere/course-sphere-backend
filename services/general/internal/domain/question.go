package domain

import "github.com/google/uuid"

type QuestionCriterion struct {
	Criterion string
	Score     int32
}

type Question struct {
	ID              uuid.UUID
	Question        string
	Position        float64
	PossibleAnswers []string
	Criteria        []QuestionCriterion
}

type CreateQuestionData struct {
	Question        string
	PossibleAnswers []string
	Criteria        []QuestionCriterion
}

type UpdateQuestionData struct {
	Question        *string
	Position        *float64
	PossibleAnswers *[]string
	Criteria        *[]QuestionCriterion
}
