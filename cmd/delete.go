package cmd

import (
	p "aws-s3-siggy/presigner"
	"context"
	"time"

	"github.com/spf13/cobra"
)

type DeleteCmdOptions struct {
	bucketName string
	objectKey  string
}

func NewDeleteCmd(opts *CmdOptions) *cobra.Command {
	deleteOpts := &DeleteCmdOptions{}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete object from s3.",
		RunE: func(cmd *cobra.Command, args []string) error {
			presigner := p.NewPresigner(opts.PresignClient)
			err := presigner.DeleteObject(context.TODO(), deleteOpts.bucketName, deleteOpts.objectKey, time.Duration(2*time.Hour))
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&deleteOpts.bucketName, "bucketName", "b", "", "Bucket name")
	cmd.MarkFlagRequired("bucketName")

	cmd.Flags().StringVarP(&deleteOpts.objectKey, "objectKey", "k", "", "Object key")
	cmd.MarkFlagRequired("objectKey")

	return cmd
}
