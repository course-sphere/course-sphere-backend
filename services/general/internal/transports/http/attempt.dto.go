package http

import (
	"time"

	"github.com/google/uuid"

	sharedDomain "github.com/course-sphere/course-sphere-backend/shared/domain"
)

type Attempt struct {
	ID         uuid.UUID         `json:"id" format:"uuid" description:"Unique attempt identifier (UUID)" example:"550e8400-e29b-41d4-a716-446655440000"`
	MaterialID uuid.UUID         `json:"material_id" format:"uuid" description:"ID of the material attempted" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Student    sharedDomain.User `json:"student" description:"Student who made the attempt"`
	Score      *int32            `json:"score,omitempty" description:"Score achieved in the attempt" example:"85"`
	CreatedAt  time.Time         `json:"created_at" format:"date-time" description:"Attempt creation timestamp" example:"2026-03-08T22:00:00Z"`
}

type AttemptDetail struct {
	AttemptID  uuid.UUID `json:"attempt_id" format:"uuid" description:"ID of the attempt" example:"550e8400-e29b-41d4-a716-446655440000"`
	QuestionID uuid.UUID `json:"question_id" format:"uuid" description:"ID of the question" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Answer     string    `json:"answer" description:"Submitted answer text" example:"B"`
}

type CreateAttemptDetailData struct {
	QuestionID uuid.UUID `json:"question_id" format:"uuid" description:"ID of the question" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Answer     string    `json:"answer" description:"Answer text submitted by the student" example:"B"`
}

type UpdateAttemptData struct {
	Score int32 `json:"score"`
}
