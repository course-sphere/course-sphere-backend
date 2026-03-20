package ports

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/services/payment/internal/domain"
	"github.com/google/uuid"
)

type WalletRepository interface {
	Create(ctx context.Context, userID uuid.UUID) error
	GetByUser(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error)
	Update(ctx context.Context, id uuid.UUID, amount int64) error
}

type HistoryRepository interface {
	Create(ctx context.Context, walletID uuid.UUID, data domain.CreateHistoryData) error
	GetByWallet(ctx context.Context, walletID uuid.UUID) ([]domain.History, error)
}
