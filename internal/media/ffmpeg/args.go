// Package ffmpeg provides argument builders for FFmpeg commands.
//
// Business logic must never construct FFmpeg arguments directly.
// All command generation goes through this package.
package ffmpeg

import (
	"fmt"
	"path/filepath"
	"time"
)

// ClipArgs holds parameters for generating a clip command.
type ClipArgs struct {
	InputFile string
	StartTime time.Duration
	EndTime   time.Duration
	OutputDir string
	Encoder   string // "copy", "h264_amf", "h264_nvenc", "h264_qsv", "libx264"
}

// ExportArgs holds parameters for a full export command.
type ExportArgs struct {
	InputFile  string
	OutputDir  string
	Encoder    string
	Format     string
	Width      int
	Height     int
	StartTime  time.Duration
	EndTime    time.Duration
	Bitrate    string
}

// BuildClipArgs constructs FFmpeg arguments for clip extraction.
// Returns the arguments and the output file path.
func BuildClipArgs(args ClipArgs) (argv []string, outputPath string) {
	outputPath = generateOutputPath(args.InputFile, args.OutputDir, "clip", args.Encoder)

	argv = append(argv,
		"-i", args.InputFile,
		"-ss", formatTimestamp(args.StartTime),
		"-to", formatTimestamp(args.EndTime),
	)

	if args.Encoder == "copy" {
		argv = append(argv, "-c", "copy")
	} else {
		argv = append(argv, "-c:v", args.Encoder)
		// Copy audio stream when re-encoding video
		argv = append(argv, "-c:a", "copy")
	}

	// Avoid overwriting existing files
	argv = append(argv, "-n")

	argv = append(argv, outputPath)
	return
}

// BuildExportArgs constructs FFmpeg arguments for full export.
func BuildExportArgs(args ExportArgs) (argv []string, outputPath string) {
	outputPath = generateOutputPath(args.InputFile, args.OutputDir, "export", args.Encoder)

	argv = append(argv, "-i", args.InputFile)

	if args.StartTime > 0 || args.EndTime > 0 {
		if args.StartTime > 0 {
			argv = append(argv, "-ss", formatTimestamp(args.StartTime))
		}
		if args.EndTime > 0 {
			argv = append(argv, "-to", formatTimestamp(args.EndTime))
		}
	}

	if args.Encoder == "copy" {
		argv = append(argv, "-c", "copy")
	} else {
		argv = append(argv, "-c:v", args.Encoder)
		argv = append(argv, "-c:a", "copy")
	}

	if args.Width > 0 && args.Height > 0 {
		argv = append(argv, "-vf", fmt.Sprintf("scale=%d:%d", args.Width, args.Height))
	}

	if args.Bitrate != "" && args.Encoder != "copy" {
		argv = append(argv, "-b:v", args.Bitrate)
	}

	argv = append(argv, "-n", outputPath)
	return
}

// generateOutputPath creates a unique output file path.
func generateOutputPath(inputFile, outputDir, suffix, encoder string) string {
	ext := ".mp4"
	if encoder == "copy" {
		ext = filepath.Ext(inputFile)
		if ext == "" {
			ext = ".mp4"
		}
	}

	base := filepath.Base(inputFile)
	name := base[:len(base)-len(filepath.Ext(base))]
	filename := fmt.Sprintf("%s_%s%s", name, suffix, ext)

	return filepath.Join(outputDir, filename)
}

// formatTimestamp converts a duration to FFmpeg timestamp format (HH:MM:SS.mmm).
func formatTimestamp(d time.Duration) string {
	total := d.Seconds()
	h := int(total) / 3600
	m := (int(total) % 3600) / 60
	s := int(total) % 60
	ms := int((total - float64(int(total))) * 1000)
	return fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)
}
