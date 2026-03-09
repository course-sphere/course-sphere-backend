package http

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
)

func (s *Server) CreateAttemptDetails(c fuego.ContextWithBody[[]CreateAttemptDetailData]) (any, error) {
	ctx := c.Context()

	id, err := uuid.Parse(c.PathParam("id"))
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
	var data []domain.CreateAttemptDetailData
	copier.Copy(&data, &raw)

	err = s.Attempt.CreateDetails(ctx, id, data)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to create attempt",
		}
	}

	return nil, nil
}

func (s *Server) GetAttemptDetails(c fuego.ContextNoBody) ([]AttemptDetail, error) {
	ctx := c.Context()

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := s.Attempt.GetDetails(ctx, id)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to get attempt details",
		}
	}
	var details []AttemptDetail
	copier.Copy(&details, &raw)

	return details, nil
}
