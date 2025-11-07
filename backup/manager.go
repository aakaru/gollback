package backup

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Manager struct {
	backupDir string
}

func NewManager(backupDir string) *Manager {
	return &Manager{
		backupDir: backupDir,
	}
}

func (m *Manager) EnsureBackupDir() error {
	if _, err := os.Stat(m.backupDir); os.IsNotExist(err) {
		err := os.MkdirAll(m.backupDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create backup directory: %v", err)
		}
		fmt.Printf("📁 Created backup directory: %s\n", m.backupDir)
	}

	return nil
}

func (m *Manager) SaveWorkflow(workflow map[string]interface{}, workflowName string) error {
	timestamp := time.Now().Format("2006-01-02_15-04-05")

	safeName := sanitizeFilename(workflowName)
	filename := fmt.Sprintf("%s_%s.json", safeName, timestamp)
	filePath := filepath.Join(m.backupDir, filename)

	jsonData, err := json.MarshalIndent(workflow, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal workflow: %v", err)
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	fmt.Printf("💾 Saved: %s\n", filename)
	return nil
}

func sanitizeFilename(name string) string {
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
