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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workflow backups",
	Long:  `Display all backup Files with their timestamps and sizes.`,
	Run: func(cmd *cobra.Command, args []string) {
		runList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList() {
	fmt.Println("📋 Listing all backups...")
	fmt.Println()

	cfg, err := config.LoadConfig()
	if err != nil {
		color.Red("❌ Configuration error:")
		fmt.Printf("   %v\n", err)
		return
	}

	if _, err := os.Stat(cfg.BackupDir); os.IsNotExist(err) {
		fmt.Println("⚠️  No backups found. Run 'gollback backup' first.")
		return
	}

	files, err := os.ReadDir(cfg.BackupDir)
	if err != nil {
		color.Red("❌ Error reading backups:")
		fmt.Printf("   %v\n", err)
		return
	}

	type FileInfo struct {
		Name string
		Path string
		Size int64
	}

	var backupFiles []FileInfo
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			info, err := file.Info()
			if err != nil {
				continue
			}
			backupFiles = append(backupFiles, FileInfo{
				Name: file.Name(),
				Path: filepath.Join(cfg.BackupDir, file.Name()),
				Size: info.Size(),
			})
		}
	}

	if len(backupFiles) == 0 {
		fmt.Println("⚠️  No backup files found.")
		return
	}

	sort.Slice(backupFiles, func(i, j int) bool {
		return backupFiles[i].Name > backupFiles[j].Name
	})

	fmt.Printf("Found %d backup file(s):\n\n", len(backupFiles))
	for i, file := range backupFiles {
		sizeKB := float64(file.Size) / 1024
		fmt.Printf("%d. %s\n", i+1, file.Name)
		fmt.Printf("   Size: %.2f KB\n", sizeKB)
		fmt.Printf("   Path: %s\n\n", file.Path)
	}
}
