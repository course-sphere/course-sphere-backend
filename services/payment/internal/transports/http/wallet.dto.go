package http

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID      uuid.UUID `json:"id" format:"uuid" description:"Unique wallet identifier" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID  uuid.UUID `json:"user_id" format:"uuid" description:"ID of the wallet owner" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Balance int64     `json:"balance" description:"Current wallet balance (in smallest currency unit)" example:"100000"`
}

type History struct {
	ID        uuid.UUID `json:"id" format:"uuid" description:"Unique transaction history identifier" example:"550e8400-e29b-41d4-a716-446655440000"`
	Amount    int64     `json:"amount" description:"Transaction amount (positive for credit, negative for debit)" example:"50000"`
	Detail    string    `json:"detail" description:"Description of the transaction" example:"Course purchase"`
	CreatedAt time.Time `json:"created_at" format:"date-time" description:"Timestamp when the transaction was created" example:"2026-03-08T22:00:00Z"`
}

type CreatePaymentLinkData struct {
	Amount      int64  `json:"amount" description:"Amount to create payment link for (in smallest currency unit)" example:"50000"`
	Description string `json:"description" example:"Buy course"`
}

type WithdrawData struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}
