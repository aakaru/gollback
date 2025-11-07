package config

func New() *Config {
	return &Config{
		N8nURL:    "http://localhost:5678/api/v1",
		APIKey:    "",
		BackupDir: "./backups",
	}
}
