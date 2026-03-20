package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jinzhu/copier"

	"github.com/course-sphere/course-sphere-backend/services/payment/internal/adapters/repo/database"
	"github.com/course-sphere/course-sphere-backend/services/payment/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/payment/internal/ports"
)

type WalletDatabase struct {
	Pool *pgxpool.Pool
}

var _ ports.WalletRepository = &WalletDatabase{}

func (db *WalletDatabase) Create(ctx context.Context, userID uuid.UUID) error {
	inner := database.New(db.Pool)

	return inner.CreateWallet(ctx, userID)
}

func (db *WalletDatabase) GetByUser(ctx context.Context, userID uuid.UUID) (*domain.Wallet, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetWalletByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var wallet domain.Wallet
	copier.Copy(&wallet, &raw)

	return &wallet, nil
}

func (db *WalletDatabase) Update(ctx context.Context, id uuid.UUID, amount int64, detail string) error {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	inner := database.New(db.Pool).WithTx(tx)

	err = inner.CreateHistory(ctx, database.CreateHistoryParams{
		WalletID: id,
		Amount:   amount,
		Detail:   detail,
	})
	if err != nil {
		return err
	}

	err = inner.UpdateWallet(ctx, database.UpdateWalletParams{
		ID:     id,
		Amount: amount,
	})
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
