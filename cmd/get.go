package cmd

import (
	p "aws-s3-siggy/presigner"
	"context"
	"time"

	"github.com/spf13/cobra"
)

type GetCmdOptions struct {
	bucketName string
	objectKey  string
}

func NewGetCmd(opts *CmdOptions) *cobra.Command {
	getOpts := &GetCmdOptions{}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Create download URL (GetObject).",
		RunE: func(cmd *cobra.Command, args []string) error {
			presigner := p.NewPresigner(opts.PresignClient)
			err := presigner.GetObject(context.TODO(), getOpts.bucketName, getOpts.objectKey, time.Duration(2*time.Hour))
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&getOpts.bucketName, "bucketName", "b", "", "Bucket name (required)")
	cmd.MarkFlagRequired("bucketName")

	cmd.Flags().StringVarP(&getOpts.objectKey, "objectKey", "k", "", "Object key (required)")
	cmd.MarkFlagRequired("objectKey")

	return cmd
}
