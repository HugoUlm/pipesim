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
)

func init() {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run a pipeline",
		Run: func(cmd *cobra.Command, args []string) {
			p, err := pipeline.Parse(file)
			if err != nil {
				fmt.Println("Failed to parse pipeline:", err)
				os.Exit(1)
			}
			pipeline.Run(p, dryRun)
		},
	}
	runCmd.Flags().StringVarP(&file, "file", "f", "", "Path to pipeline.yml (required)")
	runCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print steps without executing")
	runCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(runCmd)
}
