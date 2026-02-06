package token

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type Parser struct {
	jwkEndpoint string
}

func NewParser(authService string) Parser {
	return Parser{
		jwkEndpoint: fmt.Sprintf("%s/jwks", authService),
	}
}

func (p *Parser) Parse(rawToken string, ctx context.Context) (jwt.Token, error) {
	keyset, err := jwk.Fetch(ctx, p.jwkEndpoint)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseString(rawToken, jwt.WithKeySet(keyset))
	if err != nil {
		return nil, err
	}

	return token, nil
}
