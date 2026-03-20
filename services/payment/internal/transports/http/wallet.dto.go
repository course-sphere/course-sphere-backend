package http

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
	ID        uuid.UUID
	Amount    int64
	Detail    string
	CreatedAt time.Time
}
