package presigner

import (
	"context"
	"errors"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type MockPresignClient struct{}

var _ PresignClient = &MockPresignClient{}

func (m *MockPresignClient) PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return &v4.PresignedHTTPRequest{
		URL: "https://example.com/put",
	}, nil
}

func (m *MockPresignClient) PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return &v4.PresignedHTTPRequest{
		URL: "https://example.com/get",
	}, nil
}

func (m *MockPresignClient) PresignDeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return &v4.PresignedHTTPRequest{
		URL: "https://example.com/delete",
	}, nil
}

type ErrorMockPresignClient struct{}

var _ PresignClient = &ErrorMockPresignClient{}

func (m *ErrorMockPresignClient) PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return nil, errors.New("mock error")
}

func (m *ErrorMockPresignClient) PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return nil, errors.New("mock error")
}

func (m *ErrorMockPresignClient) PresignDeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFuncs ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	return nil, errors.New("mock error")
}
