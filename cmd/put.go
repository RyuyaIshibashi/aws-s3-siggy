package cmd

import (
	p "aws-s3-siggy/presigner"
	"context"
	"time"

	"github.com/spf13/cobra"
)

type PutCmdOptions struct {
	bucketName string
	objectKey  string
}

func NewPutCmd(opts *CmdOptions) *cobra.Command {
	putOpts := &PutCmdOptions{}

	cmd := &cobra.Command{
		Use:   "put",
		Short: "Create upload URL (PutObject).",
		RunE: func(cmd *cobra.Command, args []string) error {
			presigner := p.NewPresigner(opts.PresignClient)
			err := presigner.PutObject(context.TODO(), putOpts.bucketName, putOpts.objectKey, time.Duration(2*time.Hour))
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&putOpts.bucketName, "bucketName", "b", "", "Bucket name (required)")
	cmd.MarkFlagRequired("bucketName")

	cmd.Flags().StringVarP(&putOpts.objectKey, "objectKey", "k", "", "Object key (required)")
	cmd.MarkFlagRequired("objectKey")

	return cmd
}
