package external

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

type AuthService struct {
	BaseUrl string
}

func (s *AuthService) MustGetJwks(ctx context.Context) jwk.Set {
	jwkEndpoint := fmt.Sprintf("%s/jwks", s.BaseUrl)

	jwks, err := jwk.Fetch(context.Background(), jwkEndpoint)

	if err != nil {
		panic(fmt.Sprintf("fatal: no auth service detected at %s", jwkEndpoint))
	}

	return jwks
}
