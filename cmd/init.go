package cmd

import (
	"bufio"
	"fmt"
	"github.com/aakaru/gollback/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize gollback configuration",
	Long:  `Set up your n8n connection and backup preferences.`,
	Run: func(cmd *cobra.Command, args []string) {
		runInit()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit() {
	color.Cyan("Gollback COnfiguration Setup\n")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("n8n URL [http://localhost:5678/api/v1](http://localhost:5678/api/v1): ")
	n8nURL, _ := reader.ReadString('\n')
	n8nURL = strings.TrimSpace(n8nURL)
	if n8nURL == "" {
		n8nURL = "http://localhost:5678/api/v1"
	}

	color.Yellow("\n Get your API key from n8n:")
	fmt.Println("	1.Open n8n in browser")
	fmt.Println("	2. Settings -> n8n API")
	fmt.Println("	Create an API key")
	fmt.Println()
	fmt.Println("Enter your API key: ")
	apiKey, _ := reader.ReadString('\n')
	apiKey = strings.TrimSpace(apiKey)

	if apiKey == "" {
		color.Red("❌ API key is required!")
		return
	}

	fmt.Println("\nBackup directory [./backups]: ")
	backupDir, _ := reader.ReadString('\n')
	backupDir = strings.TrimSpace(backupDir)
	if backupDir == "" {
		backupDir = "./backups"
	}
	fmt.Println()
	color.Green("💾 Saving configuration..")
	err := config.SaveConfig(apiKey, n8nURL, backupDir)
	if err != nil {
		color.Red("❌ Error saving config: %v", err)
		return
	}

	fmt.Println()
	color.Green("✅ Configuration complete!")
	fmt.Println()
	fmt.Println("You can now run:")
	color.Cyan("	gollback backup")
	color.Cyan("	gollback list")
	color.Cyan("	gollback restore <file>")
	fmt.Println()
	color.Yellow("⚠️ Keep your .gollback file secure - it contains your API key!")
}
