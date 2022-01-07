package cmd

import (
	"github.com/mskelton/bk/cmd/pipeline"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bk <command> <subcommand> [flags]",
		Short: "Buildkite CLI",
		Long:  "Work effortlessly with Buildkite from the command line.",
	}

	cmd.AddCommand(pipeline.NewCmdPipeline())
	cmd.PersistentFlags().String("config", "", "config file (default is $HOME/.config/bk/config.yml)")

	return cmd
}
