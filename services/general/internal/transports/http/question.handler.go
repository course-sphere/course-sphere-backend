package http

import (
	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func (s *Server) MoveQuestion(c fuego.ContextWithBody[MoveQuestionData]) (any, error) {
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

	err = s.Question.Move(ctx, id, body.PrevID, body.NextID)

	return nil, err
}

func (s *Server) UpdateQuestion(c fuego.ContextWithBody[UpdateQuestionData]) (any, error) {
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
	var data domain.UpdateQuestionData
	copier.Copy(&data, &raw)

	err = s.Question.Update(ctx, id, data)

	return nil, err
}
