package token_parser

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type Parser struct {
	keyset jwk.Set
}

func New(authService string) (*Parser, error) {
	jwkEndpoint := fmt.Sprintf("%s/jwks", authService)
	keyset, err := jwk.Fetch(context.Background(), jwkEndpoint)
	if err != nil {
		return nil, err
	}

	return &Parser{keyset}, nil
}

func (p *Parser) Parse(rawToken string) (jwt.Token, error) {
	token, err := jwt.ParseString(rawToken, jwt.WithKeySet(p.keyset))
	if err != nil {
		return nil, err
	}

	return token, nil
}
