package token

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type Parser struct {
	keyset jwk.Set
}

func NewParser(authService string) Parser {
	jwkEndpoint := fmt.Sprintf("%s/jwks", authService)
	keyset, err := jwk.Fetch(context.Background(), jwkEndpoint)
	if err != nil {
		panic(fmt.Sprintf("Invalid auth service: %s", err))
	}

	return Parser{keyset}
}

func (p *Parser) Parse(rawToken string) (jwt.Token, error) {
	token, err := jwt.ParseString(rawToken, jwt.WithKeySet(p.keyset))
	if err != nil {
		return nil, err
	}

	return token, nil
}
