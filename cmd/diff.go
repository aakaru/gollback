package cmd

import (
	"fmt"
	"github.com/aakaru/gollback/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var diffCmd = &cobra.Command{
	Use:   "diff [workflow-name]",
	Short: "Compare the latest two versions of a workflow",
	Long:  `Automatically find and compare the two most recent backups of a workflow.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runDiff(args[0])
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

func runDiff(workflowName string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		color.Red("❌ Configuration error:")
		fmt.Printf("   %v\n", err)
		return
	}

	files, err := findWorkflowBackups(cfg.BackupDir, workflowName)
	if err != nil {
		color.Red("❌ Error:")
		fmt.Printf("   %v\n", err)
		return
	}

	if len(files) < 2 {
		color.Yellow("⚠️  Need at least 2 backups to compare. Found %d backup(s).", len(files))
		fmt.Println()
		fmt.Println("Tip: Run 'gollback backup' multiple times to create versions")
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i] > files[j]
	})

	latest := filepath.Join(cfg.BackupDir, files[0])
	previous := filepath.Join(cfg.BackupDir, files[1])

	color.Green("Comparing latest two versions of: %s", workflowName)
	fmt.Println()
	fmt.Printf("Latest:   %s\n", files[0])
	fmt.Printf("Previous: %s\n", files[1])
	fmt.Println()

	runCompare(latest, previous)
}

func findWorkflowBackups(backupDir, workflowName string) ([]string, error) {
	safeName := sanitizeWorkflowName(workflowName)

	files, err := os.ReadDir(backupDir)
	if err != nil {
		return nil, err
	}

	var matches []string
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), safeName) && strings.HasSuffix(file.Name(), ".json") {
			matches = append(matches, file.Name())
		}
	}

	return matches, nil
}

func sanitizeWorkflowName(name string) string {
	result := ""
	for _, char := range name {
		if (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') {
			result += string(char)
		} else if char == ' ' || char == '-' {
			result += "_"
		}
	}
	if result == "" {
		result = "workflow"
	}
	return result
}
