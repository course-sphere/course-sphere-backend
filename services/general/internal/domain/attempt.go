package domain

import (
	"time"

	"github.com/google/uuid"
)

type Attempt struct {
	ID         uuid.UUID
	MaterialID uuid.UUID
	StudentID  uuid.UUID
	Score      *int32
	CreatedAt  time.Time
}

type AttemptDetail struct {
	AttemptID  uuid.UUID
	QuestionID uuid.UUID
	Answer     string
}

type CreateAttemptDetailData struct {
	QuestionID uuid.UUID
	Answer     string
}
