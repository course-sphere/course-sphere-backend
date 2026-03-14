package usecase

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/storage/internal/ports"
)

type Storage struct {
	Inner ports.Storage
}

func (s *Storage) Upload(ctx context.Context, courseID uuid.UUID, data domain.UploadFileData) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	var key string
	if courseID == uuid.Nil {
		key = fmt.Sprintf("%s-%s", id, data.Filename)
	} else {
		key = fmt.Sprintf("%s/%s-%s", courseID, id, data.Filename)
	}

	err = s.Inner.Upload(ctx, key, data)
	if err != nil {
		return "", err
	}

	return key, nil
}

func (s *Storage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	return s.Inner.Get(ctx, key)
}
