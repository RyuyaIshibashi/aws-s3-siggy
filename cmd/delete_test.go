package cmd

import (
	"bytes"
	"errors"
	"testing"

	"aws-s3-siggy/presigner/presignerfakes"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCmd(t *testing.T) {
	tests := []struct {
		name        string
		stub        func(fake *presignerfakes.FakePresignClient)
		wantErr     bool
		expectedErr string
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
			wantErr:     true,
			expectedErr: "mock error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &presignerfakes.FakePresignClient{}
			tt.stub(fake)
			opts := &CmdOptions{PresignClient: fake}
			err := runDeleteCmd(t, []string{"--bucketName", "test-bucket", "--objectKey", "test-key"}, opts)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != "" {
					assert.Contains(t, err.Error(), tt.expectedErr)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteCmdFlags(t *testing.T) {
	fake := &presignerfakes.FakePresignClient{}
	opts := &CmdOptions{PresignClient: fake}
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
