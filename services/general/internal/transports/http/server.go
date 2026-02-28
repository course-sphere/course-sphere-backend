package http

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"

	"github.com/course-sphere/course-sphere-backend/pkg/middleware"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/config"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/usecase"
	"github.com/course-sphere/course-sphere-backend/shared/ports"
)

func NewServer(
	cfg *config.Config,
	course *usecase.Course,
	s3Client *s3.Client,
	authClient ports.AuthClient,
	userClient ports.UserClient,
) *fuego.Server {
	handler := Handler{
		Config:     cfg,
		Course:     course,
		Presigner:  s3.NewPresignClient(s3Client),
		AuthClient: authClient,
		UserClient: userClient,
	}

	s := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf(":%d", cfg.Port)),
		fuego.WithGlobalMiddlewares(middleware.Cors(cfg.CorsOrigin)),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				UIHandler:            handler.OpenAPI,
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
	handler.RegisterRoutes(s)

	return s
}
