// Package gpu detects available hardware encoders on the system.
//
// Detection runs once during application startup and results are cached.
// The application must continue to function even when no GPU is available.
package gpu

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Known FFmpeg installation paths on Windows.
var gpuSearchPaths = []string{
	`C:\ffmpeg-8.0.1\ffmpeg-2026-01-26-git-fe0813d6e2-full_build\bin`,
	`C:\ffmpeg\bin`,
	`C:\Program Files\ffmpeg\bin`,
	`C:\Program Files (x86)\ffmpeg\bin`,
}

// findFFmpeg locates the ffmpeg binary via PATH then known directories.
func findFFmpeg() (string, error) {
	if path, err := exec.LookPath("ffmpeg"); err == nil {
		return path, nil
	}
	for _, dir := range gpuSearchPaths {
		fullPath := filepath.Join(dir, "ffmpeg.exe")
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}
	return "", fmt.Errorf("ffmpeg not found")
}

// EncoderInfo describes a detected hardware encoder.
type EncoderInfo struct {
	Name      string `json:"name"`
	Available bool   `json:"available"`
	Vendor    string `json:"vendor"`
}

// Capabilities represents the GPU encoding capabilities of the system.
type Capabilities struct {
	StreamCopy    bool          `json:"stream_copy"`
	Encoders      []EncoderInfo `json:"encoders"`
	Preferred     string        `json:"preferred"`
	GPUVendor     string        `json:"gpu_vendor"`
	GPUAvailable  bool          `json:"gpu_available"`
}

// Detector probes the system for GPU encoding capabilities.
type Detector struct {
	ffmpegPath   string
	cached       *Capabilities
}

// New creates a GPU Detector.
func New() *Detector {
	path, _ := findFFmpeg()
	return &Detector{
		ffmpegPath: path,
	}
}

// Detect runs GPU capability detection and returns the results.
// Results are cached after the first call.
func (d *Detector) Detect() *Capabilities {
	if d.cached != nil {
		return d.cached
	}

	d.cached = d.detect()
	return d.cached
}

func (d *Detector) detect() *Capabilities {
	caps := &Capabilities{
		StreamCopy:   true,
		GPUAvailable: false,
	}

	if d.ffmpegPath == "" {
		caps.Encoders = []EncoderInfo{
			{Name: "h264_amf", Available: false, Vendor: "AMD"},
			{Name: "h264_nvenc", Available: false, Vendor: "NVIDIA"},
			{Name: "h264_qsv", Available: false, Vendor: "Intel"},
		}
		caps.Preferred = "libx264"
		return caps
	}

	// Query available encoders from ffmpeg
	encoders := d.queryEncoders()

	caps.Encoders = encoders

	// Select preferred encoder
	for _, enc := range encoders {
		if enc.Available {
			caps.GPUAvailable = true
			caps.GPUVendor = enc.Vendor
			caps.Preferred = enc.Name
			break
		}
	}

	if !caps.GPUAvailable {
		caps.Preferred = "libx264"
	}

	return caps
}

// queryEncoders checks which hardware encoders are available via ffmpeg.
func (d *Detector) queryEncoders() []EncoderInfo {
	cmd := exec.Command(d.ffmpegPath, "-encoders")
	output, err := cmd.Output()
	if err != nil {
		return []EncoderInfo{
			{Name: "h264_amf", Available: false, Vendor: "AMD"},
			{Name: "h264_nvenc", Available: false, Vendor: "NVIDIA"},
			{Name: "h264_qsv", Available: false, Vendor: "Intel"},
		}
	}

	outputStr := string(output)

	candidates := []struct {
		name   string
		vendor string
		flag   string
	}{
		{name: "h264_amf", vendor: "AMD", flag: "h264_amf"},
		{name: "hevc_amf", vendor: "AMD", flag: "hevc_amf"},
		{name: "h264_nvenc", vendor: "NVIDIA", flag: "nvenc"},
		{name: "hevc_nvenc", vendor: "NVIDIA", flag: "hevc_nvenc"},
		{name: "h264_qsv", vendor: "Intel", flag: "h264_qsv"},
	}

	var encoders []EncoderInfo
	for _, c := range candidates {
		available := strings.Contains(outputStr, c.flag)
		encoders = append(encoders, EncoderInfo{
			Name:      c.name,
			Available: available,
			Vendor:    c.vendor,
		})
	}

	return encoders
}
