package external

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

type AuthService struct {
	BaseUrl string
}

func (s *AuthService) GetJwks(ctx context.Context) (jwk.Set, error) {
	jwkEndpoint := fmt.Sprintf("%s/jwks", s.BaseUrl)

	return jwk.Fetch(context.Background(), jwkEndpoint)
}
