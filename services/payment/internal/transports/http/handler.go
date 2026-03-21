package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

func (s *Server) RegisterRoutes(f *fuego.Server) {
	ctx := context.Background()

	fuego.Get(f, "/", s.Ping)

	jwks := s.AuthClient.MustGetJwks(ctx)
	tokenMiddleware := middleware.RequireToken(jwks)
	authOptions := []fuego.RouteOption{
		option.Middleware(tokenMiddleware),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	}

	fuego.Post(f, "/", s.CreatePaymentLink, authOptions...)
	fuego.Post(f, "/withdraw", s.Withdraw, authOptions...)
	fuego.Post(f, "/callback", s.PaymentCallback)

	wallet := fuego.Group(f, "/wallet")
	fuego.Get(wallet, "/", s.GetWalletByUser, authOptions...)
	fuego.Get(wallet, "/histories", s.GetWalletHistories, authOptions...)
}

func (s *Server) OpenAPI(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(fmt.Sprintf("/payment%s", specURL)),
	)
}

func (s *Server) Ping(c fuego.ContextNoBody) (string, error) {
	return "pong", nil
}
