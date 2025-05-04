package cmd

import (
	"fmt"
	"go-aws-s3-presigner/presigner"
	"go-aws-s3-presigner/s3client"
	"os"

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

	s3Client, err := s3client.NewS3Client()
	if err != nil {
		panic(err)
	}

	opts := &CmdOptions{
		PresignClient: s3Client.PresignClient,
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
