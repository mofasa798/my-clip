package domain

// DownloadProgress represents the current state of a download.
type DownloadProgress struct {
	Percentage  float64 `json:"percentage"`
	Speed       string  `json:"speed"`
	ETA         string  `json:"eta"`
	BytesLoaded int64   `json:"bytes_loaded"`
	TotalBytes  int64   `json:"total_bytes"`
}

// ExportProgress represents the current state of a media export.
type ExportProgress struct {
	Percentage float64 `json:"percentage"`
	FPS        float64 `json:"fps"`
	Speed      string  `json:"speed"`
	ETA        string  `json:"eta"`
}
