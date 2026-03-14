package ports

import (
	"context"
	"io"

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
)

type Storage interface {
	Upload(ctx context.Context, key string, data domain.UploadFileData) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}
