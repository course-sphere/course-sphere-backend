package presign

import (
	"context"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/jinzhu/copier"

	"github.com/course-sphere/course-sphere-backend/services/storage/internal/domain"
	"github.com/course-sphere/course-sphere-backend/services/storage/internal/ports"
)

type S3PresignClient struct {
	inner  *s3.PresignClient
	bucket string
}

var _ ports.PresignClient = &S3PresignClient{}

func NewS3PresignClient(
	ctx context.Context,
	endpoint string,
	bucket string,
) (*S3PresignClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg, func(opt *s3.Options) {
		opt.EndpointResolverV2 = &EndpointResolver{Endpoint: endpoint}
		opt.UsePathStyle = true
	})

	presigner := s3.NewPresignClient(s3Client)

	return &S3PresignClient{
		inner:  presigner,
		bucket: bucket,
	}, nil
}

func (p *S3PresignClient) Create(ctx context.Context, data domain.CreatePresignedRequestData) (*domain.PresignedRequest, error) {
	raw, err := p.inner.PresignPostObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(p.bucket),
		Key:    aws.String(data.FileName),
		ACL:    types.ObjectCannedACLPublicRead,
	}, func(opt *s3.PresignPostOptions) {
		opt.Expires = time.Minute * 10
		opt.Conditions = append(opt.Conditions,
			map[string]any{"acl": string(types.ObjectCannedACLPublicRead)},
			map[string]any{"Content-Type": data.ContentType},
		)
	})
	if err != nil {
		return nil, err
	}
	raw.Values["acl"] = string(types.ObjectCannedACLPublicRead)
	raw.Values["Content-Type"] = data.ContentType

	var req domain.PresignedRequest
	copier.Copy(&req, raw)

	return &req, nil
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
