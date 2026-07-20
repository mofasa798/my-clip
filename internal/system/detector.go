package system

import (
	"os/exec"
	"strings"
)

// DepStatus represents the status of a dependency.
type DepStatus struct {
	Name    string `json:"name"`
	Found   bool   `json:"found"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

// DepResult holds the results of dependency detection.
type DepResult struct {
	FFmpeg  DepStatus `json:"ffmpeg"`
	FFprobe DepStatus `json:"ffprobe"`
	YtDlp   DepStatus `json:"yt_dlp"`
	AMF     DepStatus `json:"amf"`
}

// Detector checks for external dependencies on the system.
type Detector struct{}

// NewDetector creates a new Detector.
func NewDetector() *Detector {
	return &Detector{}
}

// DetectAll runs all dependency checks and returns the results.
func (d *Detector) DetectAll() *DepResult {
	return &DepResult{
		FFmpeg:  d.detectBinary("ffmpeg", "--version"),
		FFprobe: d.detectBinary("ffprobe", "--version"),
		YtDlp:   d.detectBinary("yt-dlp", "--version"),
		AMF:     d.detectAMF(),
	}
}

// detectBinary checks if a binary exists in PATH and returns its version.
func (d *Detector) detectBinary(name string, args ...string) DepStatus {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return DepStatus{
			Name:  name,
			Found: false,
		}
	}

	version := strings.TrimSpace(string(output))
	if idx := strings.Index(version, "\n"); idx >= 0 {
		version = version[:idx]
	}

	path, _ := exec.LookPath(name)
	return DepStatus{
		Name:    name,
		Found:   true,
		Version: version,
		Path:    path,
	}
}

// detectAMF checks if FFmpeg supports the h264_amf encoder.
func (d *Detector) detectAMF() DepStatus {
	cmd := exec.Command("ffmpeg", "-encoders")
	output, err := cmd.Output()
	if err != nil {
		return DepStatus{
			Name:  "h264_amf",
			Found: false,
		}
	}

	found := strings.Contains(string(output), "h264_amf")
	version := ""
	if found {
		version = "available"
	}

	path, _ := exec.LookPath("ffmpeg")
	return DepStatus{
		Name:    "h264_amf",
		Found:   found,
		Version: version,
		Path:    path,
	}
}
