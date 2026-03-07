package http

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/lestrrat-go/jwx/v3/jwt"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

func (s *Server) CreateCourse(c fuego.ContextWithBody[CreateCourseData]) (uuid.UUID, error) {
	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	userID, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	raw, err := c.Body()
	if err != nil {
		return uuid.Nil, err
	}
	course := domain.CreateCourseData{}
	copier.Copy(&course, raw)

	id, err := s.Course.Create(c.Context(), userID, course)
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid course creation data",
		}
	}

	return id, nil
}

func (s *Server) GetCourse(c fuego.ContextNoBody) (*Course, error) {
	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := s.Course.Get(c.Context(), id)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "No course with given ID",
		}
	}

	course := Course{}
	copier.Copy(&course, raw)

	instructor, err := s.UserClient.Get(c.Context(), raw.ID)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err: err,
		}
	}
	course.Instructor = instructor.Name

	return &course, nil
}
