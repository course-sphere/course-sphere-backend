package usecase

import (
	"context"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/google/uuid"

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/storage/internal/ports"
	sharedPort "github.com/course-sphere/course-sphere-backend/shared/ports"
)

var (
	UnauthorizedErr error = fmt.Errorf("User does not owned the course")
)

type Storage struct {
	Inner         ports.Storage
	GeneralClient sharedPort.GeneralClient
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

func (s *Storage) Get(ctx context.Context, token string, key string) (io.ReadCloser, error) {
	parts := strings.SplitAfterN(key, "/", 2)
	if len(parts) == 1 {
		return s.Inner.Get(ctx, key)
	}
	if token == "" {
		return nil, UnauthorizedErr
	}

	courseID, err := uuid.Parse(parts[0])
	if err != nil {
		return nil, err
	}
	enrolledCourses, err := s.GeneralClient.GetEnrolledCourses(ctx, token)
	if err != nil {
		return nil, err
	}
	if !slices.Contains(enrolledCourses, courseID) {
		return nil, UnauthorizedErr
	}

	return s.Inner.Get(ctx, key)
}
