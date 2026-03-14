package http

import (
	"io"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
)

func (s *Server) Upload(c fuego.ContextNoBody) (string, error) {
	ctx := c.Context()

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		id = uuid.Nil
	}

	req := c.Request()

	file, header, err := req.FormFile("file")
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Could not read 'file' from form data",
		}
	}
	defer file.Close()

	key, err := s.Storage.Upload(ctx, id, domain.UploadFileData{
		Filename:    header.Filename,
		ContentType: header.Header.Get("Content-Type"),
		File:        file,
	})
	if err != nil {
		return "", fuego.InternalServerError{
			Err: err,
		}
	}

	return key, nil
}

func (s *Server) Get(c fuego.ContextNoBody) (any, error) {
	ctx := c.Context()

	key := c.PathParam("key")

	file, err := s.Storage.Get(ctx, key)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid object key",
		}
	}
	defer file.Close()

	c.Response().Header().Set("Content-Disposition", "attachment; filename="+key)
	_, err = io.Copy(c.Response(), file)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err: err,
		}
	}

	return nil, nil
}
