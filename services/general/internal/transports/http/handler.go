package http

import (
	"github.com/course-sphere/course-sphere-backend/services/general/internal/usecase"
	"github.com/go-fuego/fuego"
)

type Handler struct {
	course usecase.CourseUsecase
}

func NewHandler(course usecase.CourseUsecase) *Handler {
	return &Handler{course}
}

func (h *Handler) RegisterRoutes(s *fuego.Server) {
	fuego.Get(s, "/course/:id", h.GetCourse)
}
