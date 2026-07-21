package system

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"my-clip/internal/shared"
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

// detectBinary checks if a binary exists and returns its version.
// First tries PATH lookup, then searches known installation directories.
func (d *Detector) detectBinary(name string, args ...string) DepStatus {
	// Try PATH first
	path, pathErr := exec.LookPath(name)
	if pathErr == nil {
		return d.runBinary(path, name, args...)
	}

	// Search known directories
	binaryName := name
	if filepath.Ext(name) == "" {
		binaryName = name + ".exe"
	}
	for _, dir := range shared.FFmpegSearchPaths {
		fullPath := filepath.Join(dir, binaryName)
		if _, err := os.Stat(fullPath); err == nil {
			return d.runBinary(fullPath, name, args...)
		}
	}

	return DepStatus{Name: name, Found: false}
}

// runBinary executes a binary and returns its status.
// Captures both stdout and stderr because some wrappers (e.g. Chocolatey shims)
// output version info to stderr and exit with a non-zero code.
func (d *Detector) runBinary(fullPath, displayName string, args ...string) DepStatus {
	// First try the normal approach (cmd.Output captures stdout)
	cmd := exec.Command(fullPath, args...)
	stdout, err := cmd.Output()

	var output []byte
	if err == nil {
		output = stdout
	} else {
		// On failure, try capturing stderr too — some binaries/shims
		// write version info to stderr and exit non-zero.
		cmd2 := exec.Command(fullPath, args...)
		var stderrBuf bytes.Buffer
		cmd2.Stdout = &stderrBuf
		cmd2.Stderr = &stderrBuf
		_ = cmd2.Run()
		output = stderrBuf.Bytes()
	}

	version := strings.TrimSpace(string(output))
	if version == "" {
		// Binary exists at path but we couldn't extract a version string
		absPath, _ := filepath.Abs(fullPath)
		return DepStatus{
			Name:  displayName,
			Found: true,
			Path:  absPath,
		}
	}

	if idx := strings.Index(version, "\n"); idx >= 0 {
		version = version[:idx]
	}

	absPath, _ := filepath.Abs(fullPath)
	return DepStatus{
		Name:    displayName,
		Found:   true,
		Version: version,
		Path:    absPath,
	}
}

// detectAMF checks if FFmpeg supports the h264_amf encoder.
func (d *Detector) detectAMF() DepStatus {
	// Find ffmpeg via PATH or known paths
	ffmpegPath, _ := d.findBinary("ffmpeg")
	if ffmpegPath == "" {
		return DepStatus{Name: "h264_amf", Found: false}
	}

	cmd := exec.Command(ffmpegPath, "-encoders")
	output, err := cmd.Output()
	if err != nil {
		return DepStatus{Name: "h264_amf", Found: false}
	}

	found := strings.Contains(string(output), "h264_amf")
	version := ""
	if found {
		version = "available"
	}
	return DepStatus{
		Name:    "h264_amf",
		Found:   found,
		Version: version,
		Path:    ffmpegPath,
	}
}

// findBinary locates a binary by PATH then known directories.
func (d *Detector) findBinary(name string) (string, error) {
	if path, err := exec.LookPath(name); err == nil {
		return path, nil
	}
	binaryName := name
	if filepath.Ext(name) == "" {
		binaryName = name + ".exe"
	}
	for _, dir := range shared.FFmpegSearchPaths {
		fullPath := filepath.Join(dir, binaryName)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}
	return "", fmt.Errorf("%s not found", name)
}
