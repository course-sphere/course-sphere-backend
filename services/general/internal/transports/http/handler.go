package http

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/course-sphere/course-sphere-backend/pkg/middleware"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/usecase"
)

type Handler struct {
	course *usecase.Course
}

func NewHandler(course *usecase.Course) *Handler {
	return &Handler{course}
}

func (h *Handler) RegisterRoutes(s *fuego.Server) {
	fuego.Get(s, "/", h.Ping)

	// TODO: remove hard url
	course := fuego.Group(s, "/course",
		option.Middleware(middleware.MustToken("http://auth-service:3000")),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	)
	fuego.Get(course, "/{id}", h.GetCourse)
}

func (h *Handler) OpenAPI(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(specURL),
	)
}
