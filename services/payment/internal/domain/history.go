package domain

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID        uuid.UUID
	Amount    int64
	Detail    string
	CreatedAt time.Time
}
