package http

import "github.com/google/uuid"

type QuestionCriterion struct {
	Criterion string `json:"criterion" description:"Name of the criterion" example:"clarity"`
	Score     int32  `json:"score" description:"Score weight or maximum for this criterion" example:"5"`
}

type Question struct {
	ID              uuid.UUID           `json:"id" format:"uuid" description:"Unique question identifier (UUID)" example:"550e8400-e29b-41d4-a716-446655440000"`
	Question        string              `json:"question" description:"Question text" example:"What is the derivative of x^2?"`
	Position        float64             `json:"position" description:"Position/order of the question within the quiz or material (can be fractional)" example:"1.0"`
	PossibleAnswers []string            `json:"possible_answers,omitempty" description:"List of possible answers (for multiple-choice)"`
	Criteria        []QuestionCriterion `json:"criteria,omitempty" description:"Scoring criteria for the question (criterion name + score)"`
}

type CreateQuestionData struct {
	Question        string              `json:"question" description:"Question text" example:"What is the derivative of x^2?"`
	PossibleAnswers []string            `json:"possible_answers,omitempty" description:"Possible answers for multiple-choice questions"`
	Criteria        []QuestionCriterion `json:"criteria,omitempty" description:"Scoring criteria (criterion name + score)"`
}

type MoveQuestionData struct {
	PrevID *uuid.UUID `json:"prev_id,omitempty" format:"uuid" description:"UUID of the previous question (nil to place at beginning)" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	NextID *uuid.UUID `json:"next_id,omitempty" format:"uuid" description:"UUID of the next question (nil to place at end)" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}

type UpdateQuestionData struct {
	Question        *string             `json:"question,omitempty" description:"Updated question text"`
	Position        *float64            `json:"position,omitempty" description:"Updated position/order" example:"2.5"`
	PossibleAnswers *[]string           `json:"possible_answers,omitempty" description:"Updated possible answers"`
	Criteria        *[]QuestionCriterion `json:"criteria,omitempty" description:"Updated scoring criteria"`
}
