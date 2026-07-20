// Package ffmpeg wraps the FFmpeg binary for media processing.
//
// All FFmpeg execution is centralized here. Business logic never
// constructs FFmpeg arguments directly.
package ffmpeg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"my-clip/internal/domain"
)

// Known FFmpeg installation paths on Windows.
var searchPaths = []string{
	`C:\ffmpeg-8.0.1\ffmpeg-2026-01-26-git-fe0813d6e2-full_build\bin`,
	`C:\ffmpeg\bin`,
	`C:\Program Files\ffmpeg\bin`,
	`C:\Program Files (x86)\ffmpeg\bin`,
}

// findBinary locates a binary by PATH then known directories.
func findBinary(name string) (string, error) {
	if path, err := exec.LookPath(name); err == nil {
		return path, nil
	}
	binary := name
	if filepath.Ext(name) == "" {
		binary = name + ".exe"
	}
	for _, dir := range searchPaths {
		fullPath := filepath.Join(dir, binary)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}
	return "", fmt.Errorf("%s not found", name)
}

// progressRegex matches FFmpeg progress output:
// frame=  123 fps= 30 q=28.0 size=    1024kB time=00:00:10.50 bitrate= 800.0kbits/s speed=1.00x
var progressRegex = regexp.MustCompile(
	`time=(\d{2}:\d{2}:\d{2}\.\d{2,3})`,
)

var fpsRegex = regexp.MustCompile(`fps=\s*([\d.]+)`)
var speedRegex = regexp.MustCompile(`speed=\s*([\d.]+)x`)

// Wrapper wraps the FFmpeg binary.
type Wrapper struct {
	binaryPath string
}

// New creates a new FFmpeg wrapper.
func New() (*Wrapper, error) {
	path, err := findBinary("ffmpeg")
	if err != nil {
		return nil, fmt.Errorf("ffmpeg not found: %w", err)
	}
	return &Wrapper{binaryPath: path}, nil
}

// RunClip executes FFmpeg to create a clip from the given input file.
func (w *Wrapper) RunClip(args ClipArgs, progress func(domain.ExportProgress)) error {
	argv, _ := BuildClipArgs(args)

	// Validate input file
	if _, err := os.Stat(args.InputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file not found: %s", args.InputFile)
	}

	// Validate output directory
	outputDir := args.OutputDir
	if outputDir == "" {
		outputDir = "."
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("cannot create output directory: %w", err)
	}

	return w.execute(argv, getTotalDuration(args), progress)
}

// RunExport executes FFmpeg with the given export arguments.
func (w *Wrapper) RunExport(args ExportArgs, totalDuration time.Duration, progress func(domain.ExportProgress)) error {
	argv, _ := BuildExportArgs(args)

	if _, err := os.Stat(args.InputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file not found: %s", args.InputFile)
	}

	outputDir := args.OutputDir
	if outputDir == "" {
		outputDir = "."
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("cannot create output directory: %w", err)
	}

	return w.execute(argv, totalDuration, progress)
}

// execute runs FFmpeg with the given arguments and parses progress.
func (w *Wrapper) execute(argv []string, totalDuration time.Duration, progress func(domain.ExportProgress)) error {
	// Build command
	cmd := exec.Command(w.binaryPath, argv...)

	// Capture stderr (FFmpeg outputs progress to stderr)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	// Also capture stdout for any errors
	var stdoutBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ffmpeg: %w", err)
	}

	// Parse progress from stderr
	parseProgress(stderr, totalDuration, progress)

	// Wait for completion
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("ffmpeg failed with exit code %d: %w", cmd.ProcessState.ExitCode(), err)
	}

	return nil
}

// parseProgress reads FFmpeg stderr and calls the progress callback.
func parseProgress(stderr io.ReadCloser, totalDuration time.Duration, progress func(domain.ExportProgress)) {
	if progress == nil {
		io.Copy(io.Discard, stderr)
		return
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()

		p := domain.ExportProgress{
			Percentage: 0,
			FPS:        0,
			Speed:      "0x",
			ETA:        "",
		}

		// Parse current time
		if match := progressRegex.FindStringSubmatch(line); len(match) > 1 {
			currentTime := parseFFmpegTime(match[1])
			if totalDuration > 0 {
				p.Percentage = (float64(currentTime) / float64(totalDuration)) * 100
				if p.Percentage > 100 {
					p.Percentage = 100
				}
			}
			remaining := totalDuration - currentTime
			if remaining > 0 {
				p.ETA = remaining.Round(time.Second).String()
			}
		}

		// Parse FPS
		if match := fpsRegex.FindStringSubmatch(line); len(match) > 1 {
			p.FPS, _ = strconv.ParseFloat(match[1], 64)
		}

		// Parse speed
		if match := speedRegex.FindStringSubmatch(line); len(match) > 1 {
			p.Speed = match[1] + "x"
		}

		progress(p)
	}
}

// parseFFmpegTime converts FFmpeg timestamp (HH:MM:SS.ms) to Duration.
func parseFFmpegTime(timestamp string) time.Duration {
	parts := strings.Split(timestamp, ":")
	if len(parts) != 3 {
		return 0
	}

	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	secParts := strings.Split(parts[2], ".")
	s, _ := strconv.Atoi(secParts[0])
	ms := 0
	if len(secParts) > 1 {
		ms, _ = strconv.Atoi(secParts[1])
		// Pad to milliseconds
		switch len(secParts[1]) {
		case 2:
			ms *= 10
		case 3:
			// already ms
		}
	}

	return time.Duration(h)*time.Hour +
		time.Duration(m)*time.Minute +
		time.Duration(s)*time.Second +
		time.Duration(ms)*time.Millisecond
}

// getTotalDuration extracts total duration from clip args.
func getTotalDuration(args ClipArgs) time.Duration {
	return args.EndTime - args.StartTime
}

// ValidateTimestamps checks that clip timestamps are valid.
func ValidateTimestamps(start, end time.Duration, videoDuration time.Duration) error {
	if start < 0 {
		return fmt.Errorf("start time must be >= 0, got %v", start)
	}
	if end <= start {
		return fmt.Errorf("end time (%v) must be greater than start time (%v)", end, start)
	}
	if videoDuration > 0 && end > videoDuration {
		return fmt.Errorf("end time (%v) exceeds video duration (%v)", end, videoDuration)
	}
	return nil
}
