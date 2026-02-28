package util

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

type CustomEndpointResolver struct {
	Endpoint string
}

func (r *CustomEndpointResolver) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (smithyendpoints.Endpoint, error) {
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

func NewS3ClientWithEndpoint(ctx context.Context, endpoint string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(opt *s3.Options) {
		opt.EndpointResolverV2 = &CustomEndpointResolver{Endpoint: endpoint}
		opt.UsePathStyle = true
	})

	return client, nil
}
