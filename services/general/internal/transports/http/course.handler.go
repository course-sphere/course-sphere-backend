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
	instructorID, err := uuid.Parse(sub)
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
	data := domain.CreateCourseData{}
	copier.CopyWithOption(&data, raw, copier.Option{DeepCopy: true})

	id, err := s.Course.Create(c.Context(), instructorID, data)
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

	var course Course
	copier.CopyWithOption(&course, raw, copier.Option{DeepCopy: true})

	instructor, err := s.UserClient.Get(c.Context(), raw.InstructorID)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err: err,
		}
	}
	course.Instructor = *instructor

	return &course, nil
}

func (s *Server) GetAllCourses(c fuego.ContextNoBody) ([]Course, error) {
	raw, err := s.Course.GetAll(c.Context())
	if err != nil {
		return nil, fuego.InternalServerError{
			Err: err,
		}
	}

	var courses []Course
	copier.CopyWithOption(&courses, &raw, copier.Option{DeepCopy: true})

	for i := range courses {
		instructor, err := s.UserClient.Get(c.Context(), raw[i].InstructorID)
		if err != nil {
			return nil, fuego.InternalServerError{
				Err: err,
			}
		}
		courses[i].Instructor = *instructor
	}

	return courses, nil
}

func (s *Server) UpdateCourse(c fuego.ContextWithBody[UpdateCourseData]) (any, error) {
	ctx := c.Context()

	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	instructorID, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := c.Body()
	if err != nil {
		return uuid.Nil, err
	}
	var data domain.UpdateCourseData
	copier.CopyWithOption(&data, raw, copier.Option{DeepCopy: true})

	err = s.Course.Update(ctx, id, instructorID, data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Server) CreateMaterial(c fuego.ContextWithBody[CreateMaterialData]) (uuid.UUID, error) {
	ctx := c.Context()

	courseID, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := c.Body()
	if err != nil {
		return uuid.Nil, err
	}
	var data domain.CreateMaterialData
	copier.Copy(&data, &raw)

	id, err := s.Material.Create(ctx, courseID, data)
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid material creation data",
		}
	}

	return id, nil
}

func (s *Server) GetMaterialsByCourse(c fuego.ContextNoBody) ([]Material, error) {
	ctx := c.Context()

	courseID, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := s.Material.GetByCourse(ctx, courseID)
	if err != nil {
		return nil, err
	}
	var materials []Material
	copier.Copy(&materials, &raw)

	return materials, nil
}
