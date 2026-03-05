package http

import (
	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
	"github.com/go-fuego/fuego"
	"github.com/jinzhu/copier"
)

func (s *Server) CreatePresignedRequest(c fuego.ContextWithBody[CreatePresignedRequestData]) (*PresignedRequest, error) {
	ctx := c.Context()

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	var data domain.CreatePresignedRequestData
	copier.Copy(&data, &body)

	raw, err := s.Presign.Create(ctx, data)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid presigned creation request",
		}
	}

	var req PresignedRequest
	copier.Copy(&req, raw)

	return &req, nil
}
