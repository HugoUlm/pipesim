package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/HugoUlm/pipesim/pipeline"
)

var (
	file    string
	dryRun  bool
	project string
)

func init() {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run a pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := pipeline.Parse(file)
			if err != nil {
				fmt.Println("❌ Failed to parse pipeline:", err)
				os.Exit(1)
			}

			if len(p.Steps) == 0 {
				fmt.Println("⚠️ No steps defined in pipeline.")
				return
			}
			pipeline.Run(p, dryRun, project)
		},
	}
	runCmd.Flags().StringVarP(&file, "file", "f", "", "Path to your yml-file (required)")
	runCmd.Flags().StringVarP(&project, "project", "p", "", "Path to your project")
	runCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print steps without executing")
	runCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(runCmd)
}
