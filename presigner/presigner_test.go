package presigner

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPutObject(t *testing.T) {
	tests := []struct {
		name    string
		client  PresignClient
		wantErr bool
	}{
		{
			name:    "正常系",
			client:  &MockPresignClient{},
			wantErr: false,
		},
		{
			name:    "エラーケース",
			client:  &ErrorMockPresignClient{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presigner := NewPresigner(tt.client)
			err := presigner.PutObject(context.Background(), "test-bucket", "test-key", 1*time.Hour)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, "mock error", err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetObject(t *testing.T) {
	tests := []struct {
		name    string
		client  PresignClient
		wantErr bool
	}{
		{
			name:    "正常系",
			client:  &MockPresignClient{},
			wantErr: false,
		},
		{
			name:    "エラーケース",
			client:  &ErrorMockPresignClient{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presigner := NewPresigner(tt.client)
			err := presigner.GetObject(context.Background(), "test-bucket", "test-key", 1*time.Hour)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, "mock error", err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteObject(t *testing.T) {
	tests := []struct {
		name    string
		client  PresignClient
		wantErr bool
	}{
		{
			name:    "正常系",
			client:  &MockPresignClient{},
			wantErr: false,
		},
		{
			name:    "エラーケース",
			client:  &ErrorMockPresignClient{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presigner := NewPresigner(tt.client)
			err := presigner.DeleteObject(context.Background(), "test-bucket", "test-key", 1*time.Hour)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, "mock error", err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
