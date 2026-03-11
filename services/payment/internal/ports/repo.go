package ports

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateWallet(ctx context.Context, userID uuid.UUID) error
	GetWallet(ctx context.Context, userID uuid.UUID) (int64, error)
	DepositWallet(ctx context.Context, userID uuid.UUID, amount int64) error
	WithdrawWallet(ctx context.Context, userID uuid.UUID, amount int64) error
}
