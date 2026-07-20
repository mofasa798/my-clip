package domain

import "time"

// HistoryEntry represents a single download or export record.
type HistoryEntry struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"` // "download" or "export"
	Title     string    `json:"title"`
	Source    string    `json:"source"`
	FilePath  string    `json:"file_path"`
	FileSize  int64     `json:"file_size"`
	Duration  float64   `json:"duration"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"` // "completed", "failed"
	Error     string    `json:"error,omitempty"`
}

// HistoryStore defines the interface for history persistence.
type HistoryStore interface {
	Append(entry HistoryEntry) error
	All() ([]HistoryEntry, error)
	Clear() error
}
