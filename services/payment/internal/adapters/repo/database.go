package repo

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/services/payment/internal/adapters/repo/database"
	"github.com/course-sphere/course-sphere-backend/services/payment/internal/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jinzhu/copier"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type Database struct {
	pool *pgxpool.Pool
}

var _ ports.Repository = &Database{}

func NewDatabase(cfg *config.Config) (Database, error) {
	pgxConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return Database{}, err
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		return Database{}, nil
	}

	return Database{pool}, nil
}

func (db *Database) CreateWallet(ctx context.Context, userID uuid.UUID) error {
	inner := database.New(db.pool)

	return inner.CreateWallet(ctx, userID)
}

func (db *Database) GetWallet(ctx context.Context, userID uuid.UUID) (int64, error) {
	inner := database.New(db.pool)

	return inner.GetWallet(ctx, userID)
}

func (db *Database) DepositWallet(ctx context.Context, userID uuid.UUID, amount int64) error {
	inner := database.New(db.pool)

	return inner.DepositWallet(ctx, database.DepositWalletParams{
		UserID:        userID,
		DepositAmount: amount,
	})
}

func (db *Database) WithdrawWallet(ctx context.Context, userID uuid.UUID, amount int64) error {
	inner := database.New(db.pool)

	return inner.WithdrawWallet(ctx, database.WithdrawWalletParams{
		UserID:         userID,
		WithdrawAmount: amount,
	})
}
