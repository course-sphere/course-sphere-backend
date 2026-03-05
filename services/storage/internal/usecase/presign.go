package usecase

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/storage/internal/ports"
)

type Presign struct {
	Presigner ports.PresignClient
}

func (p *Presign) Create(ctx context.Context, data domain.CreatePresignedRequestData) (*domain.PresignedRequest, error) {
	return p.Presigner.Create(ctx, data)
}
