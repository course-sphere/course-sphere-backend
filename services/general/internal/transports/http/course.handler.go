package http

import (
	"fmt"

	"github.com/course-sphere/course-sphere-backend/pkg/middleware"
	"github.com/go-fuego/fuego"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func (h *Handler) GetCourse(c fuego.ContextNoBody) (*Course, error) {
	token := c.Value(middleware.TokenKey).(jwt.Token)
	fmt.Printf("token=%+v", token)

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := h.course.Get(c.Context(), id)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "No course with given ID",
		}
	}
	course := Course{}
	copier.Copy(&course, raw)

	return &course, nil
}
