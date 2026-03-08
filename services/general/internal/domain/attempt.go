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
