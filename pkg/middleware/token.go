package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/course-sphere/course-sphere-backend/pkg/token"
	"github.com/go-fuego/fuego"
)

const (
	authorization string = "Authorization"
	bearer        string = "Bearer "
	TokenKey      string = "token"
)

func MaybeToken(authService string) func(next http.Handler) http.Handler {
	tokenParser := token.NewParser(authService)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawToken, err := getRawToken(r)
			if err != nil {
				fuego.SendJSONError(w, nil, err)
				return
			}
			if rawToken == "" {
				next.ServeHTTP(w, r)
				return
			}

			token, err := tokenParser.Parse(rawToken)
			if err != nil {
				fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
					Err:    err,
					Detail: "Invalid authorization token",
				})
				return
			}

			ctx := context.WithValue(r.Context(), TokenKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func MustToken(authService string) func(next http.Handler) http.Handler {
	tokenParser := token.NewParser(authService)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawToken, err := getRawToken(r)
			if err != nil {
				fuego.SendJSONError(w, nil, err)
				return
			}
			if rawToken == "" {
				fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
					Detail: "Missing authorization token",
				})
				return
			}

			token, err := tokenParser.Parse(rawToken)
			if err != nil {
				fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
					Err:    err,
					Detail: "Invalid authorization token",
				})
				return
			}

			ctx := context.WithValue(r.Context(), TokenKey, token)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getRawToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get(authorization)
	if authHeader == "" {
		return "", fuego.UnauthorizedError{
			Detail: "Missing authorization header",
		}
	}

	rawToken, isBearer := strings.CutPrefix(authHeader, bearer)
	if !isBearer {
		return "", fuego.UnauthorizedError{
			Detail: "Missing authorization token",
		}
	}

	return rawToken, nil
}
