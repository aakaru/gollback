package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gollback",
	Short: "A version control system for n8n workflows",
	Long: `Gollback is a CLI tool for backing up, restoring, and managing n8n workflows.
	
It provides version control capabilities for your n8n automation workflows,
including backup, restore, and comparison features.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
