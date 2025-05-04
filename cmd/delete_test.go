package cmd

import (
	"bytes"
	"testing"

	"go-aws-s3-presigner/presigner"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteCmd(t *testing.T) {
	tests := []struct {
		name        string
		mockClient  presigner.PresignClient
		wantErr     bool
		expectedErr string
	}{
		{
			name:        "正常系",
			mockClient:  &presigner.MockPresignClient{},
			wantErr:     false,
			expectedErr: "",
		},
		{
			name:        "異常系",
			mockClient:  &presigner.ErrorMockPresignClient{},
			wantErr:     true,
			expectedErr: "mock error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &CmdOptions{
				PresignClient: tt.mockClient,
			}

			err := runDeleteCmd(t, []string{"--bucketName", "test-bucket", "--objectKey", "test-key"}, opts)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDeleteCmdFlags(t *testing.T) {
	mockClient := &presigner.MockPresignClient{}
	opts := &CmdOptions{
		PresignClient: mockClient,
	}

	// フラグの必須設定確認
	err := runDeleteCmd(t, []string{}, opts)
	assert.Error(t, err, "Expected error for missing required flags")
	assert.Contains(t, err.Error(), "required flag(s) \"bucketName\", \"objectKey\" not set")
}

func runDeleteCmd(t *testing.T, args []string, opts *CmdOptions) error {
	t.Helper()
	cmd := NewDeleteCmd(opts)
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	return cmd.Execute()
}
