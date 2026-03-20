package http

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"

	"github.com/course-sphere/course-sphere-backend/services/payment/internal/config"
	"github.com/course-sphere/course-sphere-backend/services/payment/internal/usecase"
	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

type Server struct {
	Config *config.Config
	Wallet usecase.Wallet
}

func (s *Server) Build() *fuego.Server {
	f := fuego.NewServer(
		fuego.WithBasePath("/payment"),
		fuego.WithGlobalMiddlewares(middleware.Cors(s.Config.CorsOrigin)),
		fuego.WithAddr(fmt.Sprintf(":%d", s.Config.Port)),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				UIHandler:            s.OpenAPI,
				DisableDefaultServer: true,
				DisableMessages:      true,
				Info: &openapi3.Info{
					Title:       "Payment Service",
					Description: "Payment Service",
				},
			}),
		),
	)
	s.RegisterRoutes(f)

	return f
}
