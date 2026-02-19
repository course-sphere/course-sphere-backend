package ports

import (
	"context"

	"github.com/course-sphere/course-sphere-backend/shared/domain"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwk"
)

type AuthClient interface {
	MustGetJwks(ctx context.Context) jwk.Set
}

type UserClient interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
