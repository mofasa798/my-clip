package system

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config represents the application configuration.
type Config struct {
	OutputDir        string `json:"output_dir"`
	Theme            string `json:"theme"`
	PreferredEncoder string `json:"preferred_encoder"`
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	outputDir := filepath.Join(home, "Videos", "MyClip")

	return &Config{
		OutputDir:        outputDir,
		Theme:            "dark",
		PreferredEncoder: "auto",
	}
}

// configPath is overridable for testing.
var configPath = configPathDefault

// configPathDefault returns the default path to the configuration file.
func configPathDefault() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "settings.json"
	}
	dir := filepath.Join(configDir, "my-clip")
	os.MkdirAll(dir, 0755)
	return filepath.Join(dir, "settings.json")
}

// LoadConfig loads configuration from disk.
func LoadConfig() (*Config, error) {
	path := configPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// SaveConfig saves the configuration to disk.
func SaveConfig(cfg *Config) error {
	path := configPath()
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
