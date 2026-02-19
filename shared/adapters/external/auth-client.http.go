package external

import (
	"context"
	"fmt"

	"github.com/course-sphere/course-sphere-backend/shared/ports"
	"github.com/lestrrat-go/jwx/v3/jwk"
)

type HTTPAuthClient struct {
	ProxyURL string
}

var _ ports.AuthClient = &HTTPAuthClient{}

func (c *HTTPAuthClient) MustGetJwks(ctx context.Context) jwk.Set {
	jwkEndpoint := fmt.Sprintf("%s/auth/jwks", c.ProxyURL)

	jwks, err := jwk.Fetch(context.Background(), jwkEndpoint)

	if err != nil {
		panic(fmt.Sprintf("jwk fetch error: %s", err))
	}

	return jwks
}
