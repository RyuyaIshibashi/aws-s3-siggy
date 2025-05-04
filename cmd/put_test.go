package cmd

import (
	"bytes"
	"testing"

	"aws-s3-siggy/presigner"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutCmd(t *testing.T) {
	tests := []struct {
		name        string
		mockClient  presigner.PresignClient
		wantErr     bool
		expectedErr string
	}{
		{
			name:        "success case",
			mockClient:  &presigner.MockPresignClient{},
			wantErr:     false,
			expectedErr: "",
		},
		{
			name:        "error case",
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

			err := runPutCmd(t, []string{"--bucketName", "test-bucket", "--objectKey", "test-key"}, opts)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPutCmdFlags(t *testing.T) {
	mockClient := &presigner.MockPresignClient{}
	opts := &CmdOptions{
		PresignClient: mockClient,
	}

	// Verify required flags
	err := runPutCmd(t, []string{}, opts)
	assert.Error(t, err, "Expected error for missing required flags")
	assert.Contains(t, err.Error(), "required flag(s) \"bucketName\", \"objectKey\" not set")
}

func runPutCmd(t *testing.T, args []string, opts *CmdOptions) error {
	t.Helper()
	cmd := NewPutCmd(opts)
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	return cmd.Execute()
}
