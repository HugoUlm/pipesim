package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/HugoUlm/pipesim/internal/parser"
	"github.com/HugoUlm/pipesim/internal/runner"
)

var (
	file		string
	dryRun		bool
	project		string
	useCache	bool
)

func init() {
	runCmd := &cobra.Command{
		Use:   "pipesim",
		Short: "Run a pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			commands, err := parser.Parse(file, useCache, project)
			if err != nil {
				fmt.Println("❌ Failed to parse pipeline:", err)
				os.Exit(1)
			}

			if len(commands) == 0 {
				fmt.Println("⚠️ No steps defined in pipeline.")
				return
			}
			runner.Run(commands, dryRun, useCache)
		},
	}
	runCmd.Flags().StringVarP(&file, "file", "f", "", "Path to your yml-file (required)")
	runCmd.Flags().StringVarP(&project, "project", "p", "", "Path to your project")
	runCmd.Flags().BoolVar(&useCache, "use-cache", false, "Remove downloaded packages after run")
	runCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print steps without executing")
	runCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(runCmd)
}
