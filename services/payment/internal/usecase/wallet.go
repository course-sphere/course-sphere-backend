package usecase

import (
	"context"
	"fmt"

	"github.com/course-sphere/course-sphere-backend/services/payment/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/payment/internal/ports"
	"github.com/google/uuid"
)

var (
	InvalidAmountErr error = fmt.Errorf("Invalid amount")
)

type Wallet struct {
	WalletRepo  ports.WalletRepository
	HistoryRepo ports.HistoryRepository
}

func (u *Wallet) GetByUser(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
	err := u.WalletRepo.Create(ctx, userID)
	if err != nil {
		return nil, err
	}

	return u.WalletRepo.GetByUser(ctx, userID)
}

func (u *Wallet) Deposit(ctx context.Context, id uuid.UUID, amount int64, detail string) error {
	if amount <= 0 {
		return InvalidAmountErr
	}

	return u.WalletRepo.Update(ctx, id, amount, detail)
}

func (u *Wallet) Withdraw(ctx context.Context, id uuid.UUID, amount int64, detail string) error {
	if amount >= 0 {
		return InvalidAmountErr
	}

	return u.WalletRepo.Update(ctx, id, amount, detail)
}

func (u *Wallet) GetHistories(ctx context.Context, id uuid.UUID) ([]domain.History, error) {
	return u.HistoryRepo.GetByWallet(ctx, id)
}
