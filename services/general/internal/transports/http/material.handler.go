package http

import (
	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

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
