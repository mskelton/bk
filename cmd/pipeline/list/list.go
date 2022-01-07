package list

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Archived bool
	Limit    int
	Search   string
	Web      bool
}

func NewCmdList() *cobra.Command {
	opts := &ListOptions{}

	cmd := &cobra.Command{
		Use:   "bk pipeline list",
		Short: "List and filter pipelines",
		Example: heredoc.Doc(`
			List all pipelines
			$ bk pipeline list

			List favorite pipelines
			$ bk pipeline --favorite

			List piplines using search syntax
			$ bk pipeline list --search "my-pipeline"

			Open the list of pipelines in a web browser
			$ bk pipeline list --web
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Web, "web", "w", false, "Open the browser to list the pipelines")
	cmd.Flags().IntVarP(&opts.Limit, "limit", "L", 30, "Maximum number of items to fetch")
	cmd.Flags().BoolVar(&opts.Archived, "archived", false, "Only return archived pipelines")
	cmd.Flags().StringVarP(&opts.Search, "search", "S", "", "Search pipelines with `query`")

	return cmd
}

func listRun(opts *ListOptions) error {
	httpClient, err := opts.HttpClient()
	if err != nil {
		return err
	}

	if opts.Web {
		pipelineURL := util.GenerateBuildkiteURL("pipelines")

		if opts.IO.IsStdoutTTY() {
			fmt.Fprintf(opts.IO.ErrOut, "Opening %s in your browser.\n", utils.DisplayURL(pipelineURL))
		}
		return opts.Browser.Browse(pipelineURL)
	}

	listResult, err := listPullRequests(httpClient, filters, opts.LimitResults)
	if err != nil {
		return err
	}

}
