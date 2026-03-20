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

type HistoryDatabase struct {
	Pool *pgxpool.Pool
}

var _ ports.HistoryRepository = &HistoryDatabase{}

func (db *HistoryDatabase) GetByWallet(ctx context.Context, walletID uuid.UUID) ([]domain.History, error) {
	inner := database.New(db.Pool)

	raw, err := inner.GetHistoriesByWallet(ctx, walletID)
	if err != nil {
		return nil, err
	}

	var histories []domain.History
	copier.Copy(&histories, &raw)

	return histories, nil
}
