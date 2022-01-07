package pipeline

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPipelineCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pipeline <command>",
		Short: "Manage pipelines",
		Long:  "Work with Buildkite pipelines",
		Run: func(cmd *cobra.Command, args []string) {
			listRun()
		},
	}

	cmd.AddCommand(NewCmdList())

	return cmd
}

func listRun() {
	fmt.Println("warning: this query uses the Search API which is capped at 1000 results maximum")
}
