package http

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/lestrrat-go/jwx/v3/jwt"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

func (s *Server) CreateRoadmap(c fuego.ContextWithBody[CreateRoadmapData]) (uuid.UUID, error) {
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
	data := domain.CreateRoadmapData{}
	copier.Copy(&data, raw)

	id, err := s.Roadmap.Create(c.Context(), instructorID, data)
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid roadmap creation data",
		}
	}

	return id, nil
}

func (s *Server) AddRoadmapCourse(c fuego.ContextWithBody[AddCourseData]) (any, error) {
	ctx := c.Context()

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	err = s.Roadmap.AddCourse(ctx, id, body.CourseID)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to enroll",
		}
	}

	return nil, nil
}

func (s *Server) ApplyRoadmap(c fuego.ContextNoBody) (any, error) {
	ctx := c.Context()

	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	studentID, err := uuid.Parse(sub)
	if err != nil {
		return nil, fuego.UnauthorizedError{
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

	err = s.Roadmap.Apply(ctx, id, studentID)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to enroll",
		}
	}

	return nil, nil
}

func (s *Server) GetAllRoadmap(c fuego.ContextNoBody) ([]uuid.UUID, error) {
	ctx := c.Context()

	roadmaps, err := s.Roadmap.GetAll(ctx)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to get all roadmap",
		}
	}

	return roadmaps, nil
}

func (s *Server) GetRoadmapsByStudent(c fuego.ContextNoBody) ([]uuid.UUID, error) {
	ctx := c.Context()

	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	studentID, err := uuid.Parse(sub)
	if err != nil {
		return nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	roadmaps, err := s.Roadmap.GetByStudent(ctx, studentID)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to get your roadmap",
		}
	}

	return roadmaps, nil
}

func (s *Server) GetRoadmap(c fuego.ContextNoBody) (*Roadmap, error) {
	ctx := c.Context()

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := s.Course.Get(ctx, id)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid course",
		}
	}
	var roadmap Roadmap
	copier.CopyWithOption(&roadmap, &raw, copier.Option{DeepCopy: true})

	return &roadmap, nil
}

func (s *Server) MoveRoadmap(c fuego.ContextWithBody[MoveRoadmapData]) (any, error) {
	ctx := c.Context()

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	body, err := c.Body()
	if err != nil {
		return uuid.Nil, err
	}

	err = s.Material.Move(ctx, id, body.PrevID, body.NextID)

	return nil, err
}

func (s *Server) UpdateRoadmap(c fuego.ContextWithBody[UpdateRoadmapData]) (any, error) {
	ctx := c.Context()

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
	var data domain.UpdateRoadmapData
	copier.Copy(&data, &raw)

	err = s.Roadmap.Update(ctx, id, data)

	return nil, err
}
