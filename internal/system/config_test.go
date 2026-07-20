package system

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.OutputDir == "" {
		t.Error("expected non-empty OutputDir")
	}
	if cfg.Theme != "dark" {
		t.Errorf("expected theme 'dark', got %s", cfg.Theme)
	}
	if cfg.PreferredEncoder != "auto" {
		t.Errorf("expected encoder 'auto', got %s", cfg.PreferredEncoder)
	}
}

func TestSaveAndLoadConfig(t *testing.T) {
	// Use temp dir for test
	tmpDir := t.TempDir()
	origPath := configPath
	configPath = func() string {
		return filepath.Join(tmpDir, "settings.json")
	}
	defer func() { configPath = origPath }()

	// Save
	cfg := &Config{
		OutputDir:        "C:\\Videos\\MyClip",
		Theme:            "light",
		PreferredEncoder: "libx264",
	}

	if err := SaveConfig(cfg); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(configPath()); os.IsNotExist(err) {
		t.Fatal("config file was not created")
	}

	// Load
	loaded, err := LoadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if loaded.OutputDir != cfg.OutputDir {
		t.Errorf("expected OutputDir %s, got %s", cfg.OutputDir, loaded.OutputDir)
	}
	if loaded.Theme != cfg.Theme {
		t.Errorf("expected Theme %s, got %s", cfg.Theme, loaded.Theme)
	}
	if loaded.PreferredEncoder != cfg.PreferredEncoder {
		t.Errorf("expected Encoder %s, got %s", cfg.PreferredEncoder, loaded.PreferredEncoder)
	}
}

func TestLoadConfigFileNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	origPath := configPath
	configPath = func() string {
		return filepath.Join(tmpDir, "nonexistent.json")
	}
	defer func() { configPath = origPath }()

	_, err := LoadConfig()
	if err == nil {
		t.Error("expected error for missing config file")
	}
}
