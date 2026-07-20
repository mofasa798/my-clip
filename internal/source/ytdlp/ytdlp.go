// Package ytdlp provides a shared wrapper around the yt-dlp binary.
//
// Both YouTube and Kick adapters use this wrapper to retrieve metadata
// and download videos. The wrapper handles command construction,
// execution, and output parsing.
package ytdlp

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"my-clip/internal/domain"
)

// jsRuntimeArgs are added to every yt-dlp command to enable
// JavaScript-based extraction for higher-quality video formats.
// Without this flag, YouTube falls back to the android_vr client
// which only returns 360p max.
var jsRuntimeArgs = []string{"--js-runtimes", "node"}

// ytDlpSearchPaths lists additional directories to search for yt-dlp.
// The latest version is preferred over the Chocolatey/system one.
var ytDlpSearchPaths = []string{
	filepath.Join(os.Getenv("USERPROFILE"), "go", "bin"),
	filepath.Join(os.Getenv("HOME"), "go", "bin"),
}

// ytDlpOutput represents the JSON output from yt-dlp --dump-json.
type ytDlpOutput struct {
	Title      string  `json:"title"`
	Uploader   string  `json:"uploader"`
	Duration   float64 `json:"duration"`
	Thumbnail  string  `json:"thumbnail"`
	WebpageURL string  `json:"webpage_url"`
	Extractor  string  `json:"extractor"`

	// Requested formats
	FormatID       string  `json:"format_id,omitempty"`
	Ext            string  `json:"ext,omitempty"`
	Filesize       int64   `json:"filesize,omitempty"`
	FilesizeApprox int64   `json:"filesize_approx,omitempty"`
	Width          int     `json:"width,omitempty"`
	Height         int     `json:"height,omitempty"`
	VCodec         string  `json:"vcodec,omitempty"`
	ACodec         string  `json:"acodec,omitempty"`
	TBR            float64 `json:"tbr,omitempty"`

	// Available formats list
	Formats []ytDlpFormat `json:"formats,omitempty"`
}

type ytDlpFormat struct {
	FormatID   string  `json:"format_id"`
	Ext        string  `json:"ext"`
	Filesize   int64   `json:"filesize"`
	Width      int     `json:"width"`
	Height     int     `json:"height"`
	VCodec     string  `json:"vcodec"`
	ACodec     string  `json:"acodec"`
	TBR        float64 `json:"tbr"`
	FormatNote string  `json:"format_note"`
	Resolution string  `json:"resolution"`
}

// progressLine matches yt-dlp progress output like:
// [download]  45.2% of 105.34MiB at 3.45MiB/s ETA 00:17
var progressRegex = regexp.MustCompile(`\[download\]\s+([\d.]+)%\s+(?:of\s+[~\s]*([\d.]+[KMGTPEZY]?i?B))?\s*(?:at\s+([\d.]+[KMGTPEZY]?i?B/s))?\s*(?:ETA\s+(\S+))?`)

// Wrapper wraps the yt-dlp command-line tool.
type Wrapper struct {
	binaryPath string
}

// New creates a new yt-dlp wrapper.
// It finds yt-dlp in system PATH or known installation directories.
// Directories are searched first so a user-installed (latest) copy
// takes precedence over an older Chocolatey/system one.
func New() (*Wrapper, error) {
	// Search known paths first (for user-installed latest version)
	for _, dir := range ytDlpSearchPaths {
		binary := filepath.Join(dir, "yt-dlp.exe")
		if _, err := os.Stat(binary); err == nil {
			return &Wrapper{binaryPath: binary}, nil
		}
		binary = filepath.Join(dir, "yt-dlp")
		if _, err := os.Stat(binary); err == nil {
			return &Wrapper{binaryPath: binary}, nil
		}
	}
	// Fall back to PATH lookup
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		return nil, fmt.Errorf("yt-dlp not found in PATH: %w", err)
	}
	return &Wrapper{binaryPath: path}, nil
}

// ExtractMetadata retrieves video metadata from the given URL.
func (w *Wrapper) ExtractMetadata(ctx context.Context, url string) (*domain.VideoMetadata, error) {
	args := append([]string{
		"--dump-json",
		"--no-download",
		"--no-warnings",
	}, append(jsRuntimeArgs, url)...)
	cmd := exec.CommandContext(ctx, w.binaryPath, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("yt-dlp metadata failed: %w\nstderr: %s", err, stderr.String())
	}

	var raw ytDlpOutput
	if err := json.Unmarshal(stdout.Bytes(), &raw); err != nil {
		return nil, fmt.Errorf("failed to parse yt-dlp output: %w", err)
	}

	meta := &domain.VideoMetadata{
		Source:    normalizeSourceName(raw.Extractor),
		Title:     raw.Title,
		Author:    raw.Uploader,
		Duration:  time.Duration(raw.Duration * float64(time.Second)),
		Thumbnail: raw.Thumbnail,
		URL:       raw.WebpageURL,
		Streams:   buildStreams(raw.Formats),
	}

	return meta, nil
}

// Download downloads a video from the given URL.
// The progress callback is called periodically with download progress.
func (w *Wrapper) Download(ctx context.Context, req domain.DownloadRequest, progress func(domain.DownloadProgress)) (*domain.DownloadResult, error) {
	args := append([]string{
		"--newline",
		"--no-warnings",
		"-f", req.StreamID,
		"-o", filepath.Join(req.OutputDir, "%(title)s.%(ext)s"),
	}, append(jsRuntimeArgs, req.URL)...)

	cmd := exec.CommandContext(ctx, w.binaryPath, args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start yt-dlp: %w", err)
	}

	// Parse progress from stderr
	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()
		if p := parseProgress(line); p != nil && progress != nil {
			progress(*p)
		}
	}

	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("yt-dlp download failed: %w", err)
	}

	// Build result from request data
	result := &domain.DownloadResult{
		FilePath: filepath.Join(req.OutputDir, req.Filename),
	}

	return result, nil
}

// normalizeSourceName converts yt-dlp extractor names to our source names.
func normalizeSourceName(extractor string) string {
	switch {
	case strings.Contains(strings.ToLower(extractor), "youtube"):
		return "YouTube"
	case strings.Contains(strings.ToLower(extractor), "kick"):
		return "Kick"
	default:
		return extractor
	}
}

// buildStreams converts yt-dlp format list to our generic StreamInfo list.
func buildStreams(formats []ytDlpFormat) []domain.StreamInfo {
	var streams []domain.StreamInfo
	seen := make(map[string]bool)

	for _, f := range formats {
		// Skip audio-only formats if we already have a combined format
		id := f.FormatID
		if seen[id] {
			continue
		}
		seen[id] = true

		resolution := f.Resolution
		if resolution == "" && f.Width > 0 && f.Height > 0 {
			resolution = fmt.Sprintf("%dx%d", f.Width, f.Height)
		}

		quality := f.FormatNote
		if quality == "" {
			quality = resolution
		}

		size := f.Filesize
		if size == 0 {
			size = f.Filesize
		}

		streams = append(streams, domain.StreamInfo{
			ID:         id,
			Quality:    quality,
			Resolution: resolution,
			Format:     f.Ext,
			Size:       size,
			HasAudio:   f.ACodec != "none" && f.ACodec != "",
			HasVideo:   f.VCodec != "none" && f.VCodec != "",
			Bitrate:    int(f.TBR),
			Codec:      f.VCodec,
		})
	}

	return streams
}

// parseProgress attempts to parse a yt-dlp progress line.
func parseProgress(line string) *domain.DownloadProgress {
	matches := progressRegex.FindStringSubmatch(line)
	if matches == nil {
		return nil
	}

	pct, _ := strconv.ParseFloat(matches[1], 64)

	var totalBytes int64
	if matches[2] != "" {
		totalBytes = parseSize(matches[2])
	}

	return &domain.DownloadProgress{
		Percentage: pct,
		Speed:      matches[3],
		ETA:        matches[4],
		TotalBytes: totalBytes,
	}
}

// parseSize parses a human-readable size string (e.g. "105.34MiB") to bytes.
func parseSize(s string) int64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	multipliers := map[string]int64{
		"B": 1, "KiB": 1024, "MiB": 1024 * 1024,
		"GiB": 1024 * 1024 * 1024, "TiB": 1024 * 1024 * 1024 * 1024,
	}

	var numStr string
	var unit string
	for _, u := range []string{"TiB", "GiB", "MiB", "KiB", "B"} {
		if strings.HasSuffix(s, u) {
			unit = u
			numStr = strings.TrimSuffix(s, u)
			break
		}
	}

	if unit == "" {
		return 0
	}

	numStr = strings.TrimSpace(numStr)
	val, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0
	}

	return int64(val * float64(multipliers[unit]))
}
