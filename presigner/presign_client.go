package presigner

import (
	"context"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate go tool counterfeiter . PresignClient
type PresignClient interface {
	PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignDeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignUploadPart(ctx context.Context, params *s3.UploadPartInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

func NewPresigner(presignClient PresignClient) *Presigner {
	return &Presigner{PresignClient: presignClient}
}
