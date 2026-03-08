package http

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/config"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/usecase"
	"github.com/course-sphere/course-sphere-backend/shared/ports"
	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

type Server struct {
	Config *config.Config

	Course   usecase.Course
	Material usecase.Material
	Question usecase.Question

	AuthClient ports.AuthClient
	UserClient ports.UserClient
}

func (s *Server) Build() *fuego.Server {
	f := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf(":%d", s.Config.Port)),
		fuego.WithGlobalMiddlewares(middleware.Cors(s.Config.CorsOrigin)),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				UIHandler:            s.OpenAPI,
				DisableDefaultServer: true,
				DisableMessages:      true,
				Info: &openapi3.Info{
					Title:       "General Service",
					Description: "General Service",
				},
			}),
		),
		fuego.WithSecurity(openapi3.SecuritySchemes{
			"bearerAuth": &openapi3.SecuritySchemeRef{
				Value: openapi3.NewSecurityScheme().
					WithType("http").
					WithScheme("bearer").
					WithBearerFormat("JWT").
					WithDescription("Enter your JWT token in the format: Bearer <token>"),
			},
		}),
	)
	s.RegisterRoutes(f)

	return f
}
