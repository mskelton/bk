package pipeline

import (
	"github.com/mskelton/bk/cmd/pipeline/list"
	"github.com/spf13/cobra"
)

func NewCmdPipeline() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pipeline <command> [flags]",
		Short: "Manage pipelines",
		Long:  "Work with Buildkite pipelines",
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
