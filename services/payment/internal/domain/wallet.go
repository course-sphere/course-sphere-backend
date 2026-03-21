package domain

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance int64
}

type History struct {
	ID          uuid.UUID
	Amount      int64
	Description string
	CreatedAt   time.Time
}
