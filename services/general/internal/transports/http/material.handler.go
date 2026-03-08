package http

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/lestrrat-go/jwx/v3/jwt"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

func (s *Server) CreateAttempt(c fuego.ContextWithBody[CreateAttemptData]) (uuid.UUID, error) {
	ctx := c.Context()

	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	studentID, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	body, err := c.Body()
	if err != nil {
		return uuid.Nil, err
	}

	attemptID, err := s.Attempt.Create(ctx, id, studentID, body.Score)
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid attempt",
		}
	}

	return attemptID, nil
}

func (s *Server) GetMaterialAttempts(c fuego.ContextNoBody) ([]Attempt, error) {
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

	raw, err := s.Attempt.GetByMaterial(ctx, id, studentID)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to get attempts",
		}
	}
	var attempts []Attempt
	copier.Copy(&attempts, &raw)

	return attempts, nil
}

func (s *Server) MoveMaterial(c fuego.ContextWithBody[MoveMaterialData]) (any, error) {
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

func (s *Server) UpdateMaterial(c fuego.ContextWithBody[UpdateMaterialData]) (any, error) {
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
	var data domain.UpdateMaterialData
	copier.Copy(&data, &raw)

	err = s.Material.Update(ctx, id, data)

	return nil, err
}

func (s *Server) CreateQuestion(c fuego.ContextWithBody[CreateQuestionData]) (uuid.UUID, error) {
	ctx := c.Context()

	materialID, err := uuid.Parse(c.PathParam("id"))
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
	var data domain.CreateQuestionData
	copier.Copy(&data, &raw)

	id, err := s.Question.Create(ctx, materialID, data)
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid question creation data",
		}
	}

	return id, nil
}

func (s *Server) GetQuestionsByCourse(c fuego.ContextNoBody) ([]Question, error) {
	ctx := c.Context()

	materialID, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := s.Question.GetByMaterial(ctx, materialID)
	if err != nil {
		return nil, err
	}
	var questions []Question
	copier.Copy(&questions, &raw)

	return questions, nil
}
