package ports

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
)

type PresignClient interface {
	Create(ctx context.Context, data domain.CreatePresignedRequestData) (*domain.PresignedRequest, error)
}
