package domain

import (
	"time"

	"github.com/google/uuid"
)

type Attempt struct {
	Score     int64
	IsPassed  bool
	CreatedAt time.Time
}

type AttemptDetail struct {
	QuestionID    uuid.UUID
	ChosenAnswers []string
	Scores        []int64
}
