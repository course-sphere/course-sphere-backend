package middleware

import (
	"context"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

const (
	TokenKey string = "token"
)

type Service interface {
	MustGetJwks(ctx context.Context) jwk.Set
}

func RequireToken(jwks jwk.Set) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := jwt.ParseRequest(r, jwt.WithKeySet(jwks))
			if err != nil {
				err = fuego.ErrorHandler(r.Context(), fuego.UnauthorizedError{
					Err:    err,
					Detail: "Invalid authorization token",
				})
				fuego.SendJSONError(w, nil, err)
				return
			}

			ctx := context.WithValue(r.Context(), TokenKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
