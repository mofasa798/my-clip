package system

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"my-clip/internal/domain"
)

// PresetStore persists export presets to disk.
type PresetStore struct {
	mu       sync.RWMutex
	filePath string
}

// NewPresetStore creates a preset store backed by a JSON file.
func NewPresetStore() *PresetStore {
	configDir, err := os.UserConfigDir()
	path := filepath.Join(configDir, "my-clip", "presets.json")
	if err != nil {
		path = "presets.json"
	}
	os.MkdirAll(filepath.Dir(path), 0755)
	return &PresetStore{filePath: path}
}

// Load returns all saved presets, merging with defaults.
func (s *PresetStore) Load() []domain.ExportPreset {
	s.mu.RLock()
	defer s.mu.RUnlock()

	defaults := domain.DefaultPresets()
	custom := s.loadCustom()
	return append(defaults, custom...)
}

// SaveCustom persists a custom preset.
func (s *PresetStore) SaveCustom(preset domain.ExportPreset) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	presets := s.loadCustom()
	// Replace if name exists
	for i, p := range presets {
		if p.Name == preset.Name {
			presets[i] = preset
			return s.saveCustom(presets)
		}
	}
	presets = append(presets, preset)
	return s.saveCustom(presets)
}

// DeleteCustom removes a custom preset by name.
func (s *PresetStore) DeleteCustom(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	presets := s.loadCustom()
	filtered := make([]domain.ExportPreset, 0, len(presets))
	for _, p := range presets {
		if p.Name != name {
			filtered = append(filtered, p)
		}
	}
	return s.saveCustom(filtered)
}

func (s *PresetStore) loadCustom() []domain.ExportPreset {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil
	}
	var presets []domain.ExportPreset
	json.Unmarshal(data, &presets)
	return presets
}

func (s *PresetStore) saveCustom(presets []domain.ExportPreset) error {
	data, err := json.MarshalIndent(presets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0644)
}
