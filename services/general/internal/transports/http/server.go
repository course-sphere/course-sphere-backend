package http

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"

	"github.com/course-sphere/course-sphere-backend/pkg/middleware"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/config"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/usecase"
)

func NewServer(
	cfg *config.Config,
	course *usecase.Course,
) *fuego.Server {
	handler := NewHandler(course)
	s := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf(":%d", cfg.Port)),
		fuego.WithGlobalMiddlewares(middleware.Cors(cfg.AllowOrigin)),
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
	)
	handler.RegisterRoutes(s)

	return s
}
