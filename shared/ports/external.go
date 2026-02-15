package ports

import (
	"context"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

type AuthService interface {
	MustGetJwks(ctx context.Context) jwk.Set
}
