package storage

import (
	"context"
	"io"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	smithyendpoints "github.com/aws/smithy-go/endpoints"

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/storage/internal/ports"
)

type S3Storage struct {
	inner  *s3.Client
	bucket string
}

var _ ports.Storage = &S3Storage{}

func NewS3Storage(
	ctx context.Context,
	endpoint string,
	bucket string,
) (*S3Storage, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(opt *s3.Options) {
		opt.EndpointResolverV2 = &EndpointResolver{Endpoint: endpoint}
		opt.UsePathStyle = true
	})

	return &S3Storage{
		inner:  client,
		bucket: bucket,
	}, nil
}

func (s *S3Storage) Upload(ctx context.Context, key string, data domain.UploadFileData) error {
	_, err := s.inner.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        data.File,
		ContentType: aws.String(data.ContentType),
	})
	return err
}

func (s *S3Storage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	resp, err := s.inner.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

type EndpointResolver struct {
	Endpoint string
}

func (r *EndpointResolver) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (smithyendpoints.Endpoint, error) {
	if r.Endpoint == "" {
		return s3.NewDefaultEndpointResolverV2().ResolveEndpoint(ctx, params)
	}

	u, err := url.Parse(r.Endpoint)
	if err != nil {
		return smithyendpoints.Endpoint{}, err
	}
	return smithyendpoints.Endpoint{
		URI: *u,
	}, nil
}
