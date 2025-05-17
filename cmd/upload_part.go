package cmd

import (
	p "aws-s3-siggy/presigner"
	"context"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type UploadPartCmdOptions struct {
	bucketName string
	objectKey  string
	uploadId   string
	partNumber string
}

func NewUploadPartCmd(opts *CmdOptions) *cobra.Command {
	uploadPartOpts := &UploadPartCmdOptions{}

	cmd := &cobra.Command{
		Use:   "upload_part",
		Short: "Create upload part URL (UploadPart).",
		RunE: func(cmd *cobra.Command, args []string) error {
			partNumber, err := strconv.ParseInt(uploadPartOpts.partNumber, 10, 32)
			if err != nil {
				return err
			}

			presigner := p.NewPresigner(opts.PresignClient)
			err = presigner.UploadPart(
				context.TODO(),
				uploadPartOpts.bucketName,
				uploadPartOpts.objectKey,
				uploadPartOpts.uploadId,
				int32(partNumber),
				time.Duration(2*time.Hour),
			)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&uploadPartOpts.bucketName, "bucketName", "b", "", "Bucket name (required)")
	cmd.MarkFlagRequired("bucketName")

	cmd.Flags().StringVarP(&uploadPartOpts.objectKey, "objectKey", "k", "", "Object key (required)")
	cmd.MarkFlagRequired("objectKey")

	cmd.Flags().StringVarP(&uploadPartOpts.uploadId, "uploadId", "u", "", "Upload ID (required)")
	cmd.MarkFlagRequired("uploadId")

	cmd.Flags().StringVarP(&uploadPartOpts.partNumber, "partNumber", "p", "", "Part number (required)")
	cmd.MarkFlagRequired("partNumber")

	return cmd
}
