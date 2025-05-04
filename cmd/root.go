package cmd

import (
	"context"
	"fmt"
	"go-aws-s3-presigner/presigner"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

var cfgFile string

type CmdOptions struct {
	PresignClient presigner.PresignClient
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "siggy",
		Short: "Command name argument expected.",
		Long: `Available command groups for siggy:

	  put         Put object to s3
	  get         Get object from s3
	  delete      Delete object from s3
	`,
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd-test.yaml)")

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	opts := &CmdOptions{
		PresignClient: presignClient,
	}

	cmd.AddCommand(NewPutCmd(opts))
	cmd.AddCommand(NewGetCmd(opts))
	cmd.AddCommand(NewDeleteCmd(opts))

	return cmd
}

func Execute() {
	cmd := NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
