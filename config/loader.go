package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	N8nURL    string
	APIKey    string
	BackupDir string
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("n8n_url", "http://localhost:5678/api/v1")
	viper.SetDefault("backup_dir", "./backups")

	viper.SetConfigName(".gollbackrc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %v", err)
		}
	}

	viper.SetEnvPrefix("GOLLBACK")
	viper.AutomaticEnv()

	config := &Config{
		N8nURL:    viper.GetString("n8n_url"),
		APIKey:    viper.GetString("api_key"),
		BackupDir: viper.GetString("backup_dir"),
	}
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key not configured. Set it via:\n" +
			"	1. Environment variable: export GOLLBACK_API_KEY=your-key\n" +
			"	2. Config file: Create /.gollbackrc or ./.gollbackrc\n" +
			"	3. Run: gollback init")
	}
	return config, nil
}

func SaveConfig(apiKey, n8nURL, backupDir string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(homeDir, ".gollbackrc")

	content := "# Gollback Configuration File\n"
	content += "# Generated automatically - keep this file secure!\n"
	content += "\n"
	content += "# n8n API Configuration\n"
	content += "api_key: " + apiKey + "\n"
	content += "n8n_url: " + n8nURL + "\n"
	content += "\n"
	content += "# Backup Settings\n"
	content += "backup_dir: " + backupDir + "\n"

	err = os.WriteFile(configPath, []byte(content), 0600)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Configuration saved to: %s\n", configPath)
	return nil
}
