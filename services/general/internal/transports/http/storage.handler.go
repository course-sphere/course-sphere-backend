package http

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
		ACL:         "public-read",
	}, func(opt *s3.PresignPostOptions) {
		opt.Expires = time.Minute * 10
	})
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid file",
		}
	}

	var resp PresignedURL
	copier.Copy(&resp, raw)

	return &resp, nil
}
