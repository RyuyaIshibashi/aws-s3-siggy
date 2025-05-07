package presigner_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"aws-s3-siggy/presigner"
	"aws-s3-siggy/presigner/presignerfakes"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/stretchr/testify/assert"
)

func TestPutObject(t *testing.T) {
	tests := []struct {
		name    string
		stub    func(*presignerfakes.FakePresignClient)
		wantErr bool
	}{
		{
			name: "success case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignPutObjectReturns(&v4.PresignedHTTPRequest{URL: "https://example.com/put"}, nil)
			},
			wantErr: false,
		},
		{
			name: "error case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignPutObjectReturns(nil, errors.New("mock error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &presignerfakes.FakePresignClient{}
			tt.stub(fake)
			p := presigner.NewPresigner(fake)
			err := p.PutObject(context.Background(), "test-bucket", "test-key", time.Hour)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetObject(t *testing.T) {
	tests := []struct {
		name    string
		stub    func(*presignerfakes.FakePresignClient)
		wantErr bool
	}{
		{
			name: "success case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignGetObjectReturns(&v4.PresignedHTTPRequest{URL: "https://example.com/get"}, nil)
			},
			wantErr: false,
		},
		{
			name: "error case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignGetObjectReturns(nil, errors.New("mock error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &presignerfakes.FakePresignClient{}
			tt.stub(fake)
			p := presigner.NewPresigner(fake)
			err := p.GetObject(context.Background(), "test-bucket", "test-key", time.Hour)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteObject(t *testing.T) {
	tests := []struct {
		name    string
		stub    func(*presignerfakes.FakePresignClient)
		wantErr bool
	}{
		{
			name: "success case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignDeleteObjectReturns(&v4.PresignedHTTPRequest{URL: "https://example.com/delete"}, nil)
			},
			wantErr: false,
		},
		{
			name: "error case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignDeleteObjectReturns(nil, errors.New("mock error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &presignerfakes.FakePresignClient{}
			tt.stub(fake)
			p := presigner.NewPresigner(fake)
			err := p.DeleteObject(context.Background(), "test-bucket", "test-key", time.Hour)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUploadPart(t *testing.T) {
	tests := []struct {
		name    string
		stub    func(*presignerfakes.FakePresignClient)
		wantErr bool
	}{
		{
			name: "success case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignUploadPartReturns(&v4.PresignedHTTPRequest{URL: "https://example.com/upload_part"}, nil)
			},
			wantErr: false,
		},
		{
			name: "error case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignUploadPartReturns(nil, errors.New("mock error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &presignerfakes.FakePresignClient{}
			tt.stub(fake)
			p := presigner.NewPresigner(fake)
			err := p.UploadPart(context.Background(), "test-bucket", "test-key", "test-upload-id", 1, time.Hour)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
