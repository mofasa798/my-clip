package system

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"my-clip/internal/domain"
)

// HistoryStore persists download and export history to disk.
type HistoryStore struct {
	mu      sync.RWMutex
	entries []domain.HistoryEntry
	filePath string
}

// NewHistoryStore creates a history store backed by a JSON file.
func NewHistoryStore() *HistoryStore {
	configDir, err := os.UserConfigDir()
	path := filepath.Join(configDir, "my-clip", "history.json")
	if err != nil {
		path = "history.json"
	}
	os.MkdirAll(filepath.Dir(path), 0755)

	return &HistoryStore{
		filePath: path,
	}
}

// Append adds an entry to the history and persists to disk.
func (s *HistoryStore) Append(entry domain.HistoryEntry) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.entries = append(s.entries, entry)
	return s.save()
}

// All returns all history entries.
func (s *HistoryStore) All() ([]domain.HistoryEntry, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.entries == nil {
		s.load()
	}

	result := make([]domain.HistoryEntry, len(s.entries))
	copy(result, s.entries)
	return result, nil
}

// Delete removes a single history entry by index (from end).
func (s *HistoryStore) Delete(index int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.entries == nil {
		s.load()
	}

	if index < 0 || index >= len(s.entries) {
		return nil
	}

	s.entries = append(s.entries[:index], s.entries[index+1:]...)
	return s.save()
}

// Clear removes all history entries.
func (s *HistoryStore) Clear() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.entries = nil
	return s.save()
}

func (s *HistoryStore) load() {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		s.entries = []domain.HistoryEntry{}
		return
	}
	json.Unmarshal(data, &s.entries)
	if s.entries == nil {
		s.entries = []domain.HistoryEntry{}
	}
}

func (s *HistoryStore) save() error {
	data, err := json.MarshalIndent(s.entries, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0644)
}
