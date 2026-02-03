package http

import (
	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
)

func (h *Handler) GetCourse(c fuego.ContextNoBody) (*domain.Course, error) {
	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	course, err := h.course.Get(c.Context(), id)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "No course with given ID",
		}
	}

	return course, nil
}
