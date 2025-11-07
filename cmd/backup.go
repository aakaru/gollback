package cmd

import (
	"fmt"
	"github.com/aakaru/gollback/api"
	"github.com/aakaru/gollback/backup"
	"github.com/aakaru/gollback/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup all n8n workflows",
	Long:  `Fetches all workflows from n8n and saves them as timestamped JSON files.`,
	Run: func(cmd *cobra.Command, args []string) {
		runBackup()
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}

func runBackup() {
	fmt.Println("🚀 n8n Backup Tool Starting...")
	fmt.Println()

	cfg, err := config.LoadConfig()
	if err != nil {
		color.Red("❌ Configuration error:")
		fmt.Printf("   %v\n", err)
		return
	}

	client := api.NewClient(cfg)
	backupMgr := backup.NewManager(cfg.BackupDir)

	if err := backupMgr.EnsureBackupDir(); err != nil {
		color.Red("❌ Error creating backup directory:")
		fmt.Printf("   %v\n", err)
		return
	}
	fmt.Println()

	fmt.Println("📡 Fetching workflows from n8n...")
	workflows, err := client.GetWorkflows()
	if err != nil {
		color.Red("❌ Error fetching workflows:")
		fmt.Printf("   %v\n", err)
		return
	}

	if len(workflows) == 0 {
		fmt.Println("⚠️  No workflows found to backup")
		return
	}

	fmt.Printf("✅ Found %d workflow(s)\n\n", len(workflows))
	fmt.Println("💾 Starting backup...")

	successCount := 0
	for _, wf := range workflows {
		fmt.Printf("   Backing up: %s...\n", wf.Name)

		fullWorkflow, err := client.GetWorkflowByID(wf.ID)
		if err != nil {
			fmt.Printf("   ❌ Failed: %v\n", err)
			continue
		}

		err = backupMgr.SaveWorkflow(fullWorkflow, wf.Name)
		if err != nil {
			fmt.Printf("   ❌ Failed to save: %v\n", err)
			continue
		}
		successCount++
	}

	fmt.Println()
	fmt.Println("═══════════════════════════════════════")
	fmt.Printf("✅ Backup Complete! %d/%d workflows saved\n", successCount, len(workflows))
	fmt.Printf("📂 Location: %s\n", cfg.BackupDir)
	fmt.Println("═══════════════════════════════════════")
}
