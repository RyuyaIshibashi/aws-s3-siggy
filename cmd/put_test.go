package cmd

import (
	"bytes"
	"errors"
	"testing"

	"aws-s3-siggy/presigner/presignerfakes"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/stretchr/testify/assert"
)

func TestPutCmd(t *testing.T) {
	tests := []struct {
		name        string
		stub        func(fake *presignerfakes.FakePresignClient)
		wantedErr   bool
		expectedErr string
	}{
		{
			name: "success case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignPutObjectReturns(&v4.PresignedHTTPRequest{URL: "https://example.com/put"}, nil)
			},
			wantedErr: false,
		},
		{
			name: "error case",
			stub: func(fake *presignerfakes.FakePresignClient) {
				fake.PresignPutObjectReturns(nil, errors.New("mock error"))
			},
			wantedErr:   true,
			expectedErr: "mock error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &presignerfakes.FakePresignClient{}
			tt.stub(fake)
			opts := &CmdOptions{PresignClient: fake}
			err := runPutCmd(t, []string{"--bucketName", "test-bucket", "--objectKey", "test-key"}, opts)
			if tt.wantedErr {
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

func TestPutCmdFlags(t *testing.T) {
	fake := &presignerfakes.FakePresignClient{}
	opts := &CmdOptions{PresignClient: fake}
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
