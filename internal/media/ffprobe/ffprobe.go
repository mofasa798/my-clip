// Package ffprobe wraps the FFprobe binary for reading local media metadata.
package ffprobe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"my-clip/internal/shared"
)

// findFFprobe locates the ffprobe binary.
func findFFprobe() (string, error) {
	if path, err := exec.LookPath("ffprobe"); err == nil {
		return path, nil
	}
	for _, dir := range shared.FFmpegSearchPaths {
		fullPath := filepath.Join(dir, "ffprobe.exe")
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}
	return "", fmt.Errorf("ffprobe not found")
}

// MediaInfo holds metadata about a local media file.
type MediaInfo struct {
	FilePath  string  `json:"file_path"`
	Duration  float64 `json:"duration_seconds"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Codec     string  `json:"codec"`
	Bitrate   int     `json:"bitrate"`
	FPS       float64 `json:"fps"`
	SizeBytes int64   `json:"size_bytes"`
	HasAudio  bool    `json:"has_audio"`
	HasVideo  bool    `json:"has_video"`
	Format    string  `json:"format"`
}

// ffprobeOutput represents the JSON structure from ffprobe.
type ffprobeOutput struct {
	Streams []ffprobeStream `json:"streams"`
	Format  ffprobeFormat   `json:"format"`
}

type ffprobeStream struct {
	CodecType  string `json:"codec_type"`
	CodecName  string `json:"codec_name"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Bitrate    string `json:"bitrate"`
	RFrameRate string `json:"r_frame_rate"`
}

type ffprobeFormat struct {
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	FormatName string `json:"format_name"`
	Bitrate    string `json:"bitrate"`
}

// Probe reads metadata from a local media file.
func Probe(filePath string) (*MediaInfo, error) {
	path, err := findFFprobe()
	if err != nil {
		return nil, fmt.Errorf("ffprobe not found: %w", err)
	}

	args := []string{
		"-v", "quiet",
		"-print_format", "json",
		"-show_streams",
		"-show_format",
		filePath,
	}

	cmd := exec.Command(path, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("ffprobe failed: %w\nstderr: %s", err, stderr.String())
	}

	var raw ffprobeOutput
	if err := json.Unmarshal(stdout.Bytes(), &raw); err != nil {
		return nil, fmt.Errorf("failed to parse ffprobe output: %w", err)
	}

	info := &MediaInfo{
		FilePath: filePath,
	}

	// Parse duration
	if raw.Format.Duration != "" {
		info.Duration, _ = strconv.ParseFloat(raw.Format.Duration, 64)
	}

	// Parse size
	if raw.Format.Size != "" {
		size, _ := strconv.ParseInt(raw.Format.Size, 10, 64)
		info.SizeBytes = size
	}

	info.Format = raw.Format.FormatName

	// Parse bitrate
	if raw.Format.Bitrate != "" {
		br, _ := strconv.Atoi(raw.Format.Bitrate)
		info.Bitrate = br
	}

	// Parse streams
	for _, s := range raw.Streams {
		switch s.CodecType {
		case "video":
			info.HasVideo = true
			info.Codec = s.CodecName
			info.Width = s.Width
			info.Height = s.Height
			info.FPS = parseFPS(s.RFrameRate)
			if s.Bitrate != "" {
				if br, err := strconv.Atoi(s.Bitrate); err == nil {
					info.Bitrate = br
				}
			}
		case "audio":
			info.HasAudio = true
		}
	}

	return info, nil
}

// parseFPS converts a frame rate string like "30000/1001" to float64.
func parseFPS(fpsStr string) float64 {
	if fpsStr == "" {
		return 0
	}
	parts := strings.Split(fpsStr, "/")
	if len(parts) != 2 {
		f, _ := strconv.ParseFloat(fpsStr, 64)
		return f
	}
	num, _ := strconv.ParseFloat(parts[0], 64)
	den, _ := strconv.ParseFloat(parts[1], 64)
	if den == 0 {
		return 0
	}
	return num / den
}
