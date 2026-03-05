package http

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/go-fuego/fuego"
	"github.com/jinzhu/copier"
)

func (h *Handler) CreateS3PresignedURL(c fuego.ContextWithBody[CreatePresignedURLRequest]) (*PresignedURL, error) {
	ctx := c.Context()

	req, err := c.Body()
	if err != nil {
		return nil, err
	}

	raw, err := h.Presigner.PresignPostObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(h.Config.S3Bucket),
		Key:         aws.String(req.FileName),
		ContentType: aws.String(req.ContentType),
	}, func(opt *s3.PresignPostOptions) {
		opt.Expires = time.Minute * 10
		opt.Conditions = append(opt.Conditions,
			map[string]any{"acl": string(types.ObjectCannedACLPublicRead)},
			map[string]any{"Content-Type": req.ContentType},
		)
	})
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid file",
		}
	}
	raw.Values["acl"] = string(types.ObjectCannedACLPublicRead)
	raw.Values["Content-Type"] = req.ContentType

	var resp PresignedURL
	copier.Copy(&resp, raw)

	return &resp, nil
}
