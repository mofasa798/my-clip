package domain

import "time"

// VideoMetadata represents generic video information from any source.
type VideoMetadata struct {
	Source    string        `json:"source"`
	Title     string        `json:"title"`
	Author    string        `json:"author"`
	Duration  time.Duration `json:"duration"`
	Thumbnail string        `json:"thumbnail"`
	URL       string        `json:"url"`
	Streams   []StreamInfo  `json:"streams"`
}

// StreamInfo describes an available video/audio stream.
type StreamInfo struct {
	ID          string `json:"id"`
	Quality     string `json:"quality"`
	Resolution  string `json:"resolution"`
	Format      string `json:"format"`
	Size        int64  `json:"size"`
	HasAudio    bool   `json:"has_audio"`
	HasVideo    bool   `json:"has_video"`
	Bitrate     int    `json:"bitrate"`
	Codec       string `json:"codec"`
}

// DownloadRequest contains parameters for downloading media.
type DownloadRequest struct {
	URL       string `json:"url"`
	StreamID  string `json:"stream_id"`
	OutputDir string `json:"output_dir"`
	Filename  string `json:"filename"`
}

// DownloadResult contains the result of a download operation.
type DownloadResult struct {
	FilePath string        `json:"file_path"`
	Size     int64         `json:"size"`
	Duration time.Duration `json:"duration"`
	Format   string        `json:"format"`
}

// ClipRequest represents a request to create a clip.
type ClipRequest struct {
	InputFile string        `json:"input_file"`
	StartTime time.Duration `json:"start_time"`
	EndTime   time.Duration `json:"end_time"`
	OutputDir string        `json:"output_dir"`
}

// ExportRequest represents a request to export media.
type ExportRequest struct {
	InputFile  string `json:"input_file"`
	OutputDir  string `json:"output_dir"`
	Encoder    string `json:"encoder"`
	Format     string `json:"format"`
	Preset     string `json:"preset"`
}

// EncoderOption describes an available encoder choice.
type EncoderOption struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Available bool   `json:"available"`
}
