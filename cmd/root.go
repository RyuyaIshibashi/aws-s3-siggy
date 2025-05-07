package cmd

import (
	"aws-s3-siggy/presigner"
	"aws-s3-siggy/s3client"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

type CmdOptions struct {
	PresignClient presigner.PresignClient
}

func NewCmdRoot(version, revision string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "siggy",
		Short: "Command name argument expected.",
		Long:  usage(version, revision),
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
	cmd.AddCommand(NewUploadPartCmd(opts))

	return cmd
}

func Execute(version, revision string) {
	cmd := NewCmdRoot(version, revision)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func usage(version, revision string) string {
	format := `
      _                   
  ___(_) __ _  __ _ _   _ 
 / __| |/ _' |/ _' | | | |
 |__ | | (_| | (_| | |_| |
 |___|_|___, |___, |___, |
        |___/ |___/ |___/ 
  Version: V%s-%s

Available command groups for siggy:

	  put         Create upload URL (PutObject)
	  get         Create download URL (GetObject)
	  delete      Create delete URL (DeleteObject)
	  upload_part Create upload part URL (UploadPart)

Author:
  Ryuya Ishibashi
`
	return fmt.Sprintf(format, version, revision)
}
