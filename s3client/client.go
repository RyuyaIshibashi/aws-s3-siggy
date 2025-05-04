package s3client

import (
	"aws-s3-siggy/presigner"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	PresignClient *s3.PresignClient
}

var _ presigner.PresignClient = (*s3.PresignClient)(nil)

func NewS3Client() (*S3Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("configuration error: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	return &S3Client{
		PresignClient: presignClient,
	}, nil
}
