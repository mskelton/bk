package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "bk",
  Short: "Work effortlessly with Buildkite from the command line.",
	Long: `Search pipelines and builds; start and cancel builds; and much more.
						Built with love by mskelton in Go.`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
