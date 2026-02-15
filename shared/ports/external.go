package ports

import (
	"context"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

type AuthService interface {
	GetJwks(ctx context.Context) (jwk.Set, error)
}
