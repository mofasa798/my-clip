package domain

// ExportPreset represents a saved export configuration.
type ExportPreset struct {
	Name    string `json:"name"`
	Encoder string `json:"encoder"`
	Format  string `json:"format"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Bitrate string `json:"bitrate"`
}

// DefaultPresets returns the built-in export presets.
func DefaultPresets() []ExportPreset {
	return []ExportPreset{
		{Name: "Fast (Stream Copy)", Encoder: "copy", Format: "mp4"},
		{Name: "Balanced (GPU)", Encoder: "auto", Format: "mp4"},
		{Name: "Maximum Quality", Encoder: "libx264", Format: "mp4", Bitrate: "20M"},
	}
}
