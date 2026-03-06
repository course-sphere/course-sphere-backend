package http

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/lestrrat-go/jwx/v3/jwt"

	"github.com/course-sphere/course-sphere-backend/services/general/internal/domain"
	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

func (s *Server) CreateMaterial(c fuego.ContextWithBody[CreateMaterial]) (uuid.UUID, error) {
	courseID, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

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
	material := domain.CreateMaterial{}
	copier.Copy(&material, raw)

	materialID, err := s.Course.CreateMaterial(c.Context(), courseID, userID, material)
	if err != nil {
		return uuid.Nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid course creation data",
		}
	}

	return materialID, nil
}

func (s *Server) GetMaterials(c fuego.ContextNoBody) ([]Material, error) {
	courseID, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := s.Course.GetMaterials(c.Context(), courseID)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "No course with given ID",
		}
	}

	var materials []Material
	copier.Copy(&materials, raw)

	return materials, nil
}

func (s *Server) UpdateMaterial(c fuego.ContextWithBody[UpdateMaterial]) (any, error) {
	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	userID, err := uuid.Parse(sub)
	if err != nil {
		return nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	materialID, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	raw, err := c.Body()
	if err != nil {
		return nil, err
	}
	material := domain.UpdateMaterial{}
	copier.Copy(&material, raw)

	err = s.Course.UpdateMaterial(c.Context(), materialID, userID, material)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid course update data",
		}
	}

	return nil, nil
}

func (s *Server) DeleteMaterial(c fuego.ContextNoBody) (any, error) {
	token := c.Value(middleware.TokenKey).(jwt.Token)
	sub, _ := token.Subject()
	userID, err := uuid.Parse(sub)
	if err != nil {
		return nil, fuego.UnauthorizedError{
			Err:    err,
			Detail: "Invalid token",
		}
	}

	materialID, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "ID must be UUIDv4",
		}
	}

	err = s.Course.DeleteMaterial(c.Context(), materialID, userID)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "No course with given id",
		}
	}

	return nil, nil
}
