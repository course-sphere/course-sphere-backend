package http

import (
	"net/http"

	"github.com/go-fuego/fuego"
	httpSwagger "github.com/swaggo/http-swagger/v2"

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
	fuego.Get(s, "/course/{id}", h.GetCourse)
}

func (h *Handler) Ping(c fuego.ContextNoBody) (string, error) {
	return "pong", nil
}

func (h *Handler) OpenAPI(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.BaseLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(specURL),
	)
}
