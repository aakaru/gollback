package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/aakaru/gollback/api"
	"github.com/aakaru/gollback/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var restoreCmd = &cobra.Command{
	Use:   "restore [backup-file]",
	Short: "Restore a workflow from backup",
	Long:  `Restore a workflow to n8n from a backup JSON file.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runRestore(args[0])
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}

func runRestore(filename string) {
	fmt.Printf("🔄 Restoring workflow from: %s\n\n", filename)

	cfg, err := config.LoadConfig()
	if err != nil {
		color.Red("❌ Configuration error:")
		fmt.Printf("   %v\n", err)
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		color.Red("❌ Error reading file:")
		fmt.Printf("   %v\n", err)
		return
	}

	var workflow map[string]interface{}
	if err := json.Unmarshal(data, &workflow); err != nil {
		color.Red("❌ Error parsing JSON:")
		fmt.Printf("   %v\n", err)
		return
	}

	workflowName, ok := workflow["name"].(string)
	if !ok {
		workflowName = "Unknown"
	}

	fmt.Printf("📋 Workflow: %s\n", workflowName)

	delete(workflow, "id")

	client := api.NewClient(cfg)

	fmt.Println("📤 Uploading to n8n...")
	if err := client.CreateWorkflow(workflow); err != nil {
		color.Red("❌ Error:")
		fmt.Printf("   %v\n", err)
		return
	}

	fmt.Println()
	fmt.Println("✅ Workflow restored successfully!")
}
