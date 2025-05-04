package cmd

import (
	"context"
	p "go-aws-s3-presigner/presigner"
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
		Short: "Get object from s3.",
		RunE: func(cmd *cobra.Command, args []string) error {
			presigner := p.NewPresigner(opts.PresignClient)
			err := presigner.GetObject(context.TODO(), getOpts.bucketName, getOpts.objectKey, time.Duration(2*time.Hour))
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&getOpts.bucketName, "bucketName", "b", "", "Bucket name")
	cmd.MarkFlagRequired("bucketName")

	cmd.Flags().StringVarP(&getOpts.objectKey, "objectKey", "k", "", "Object key")
	cmd.MarkFlagRequired("objectKey")

	return cmd
}
