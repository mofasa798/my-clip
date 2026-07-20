package ffmpeg

import (
	"strings"
	"testing"
	"time"
)

func TestFormatTimestamp(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{0, "00:00:00.000"},
		{time.Second, "00:00:01.000"},
		{time.Minute, "00:01:00.000"},
		{time.Hour, "01:00:00.000"},
		{90*time.Second + 500*time.Millisecond, "00:01:30.500"},
		{3661*time.Second + 123*time.Millisecond, "01:01:01.123"},
	}

	for _, tt := range tests {
		result := formatTimestamp(tt.input)
		if result != tt.expected {
			t.Errorf("formatTimestamp(%v) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}

func TestBuildClipArgsStreamCopy(t *testing.T) {
	args := ClipArgs{
		InputFile: "/tmp/test.mp4",
		StartTime: 10 * time.Second,
		EndTime:   30 * time.Second,
		OutputDir: "/tmp/output",
		Encoder:   "copy",
	}

	argv, outputPath := BuildClipArgs(args)

	if outputPath == "" {
		t.Error("expected non-empty output path")
	}
	if !strings.Contains(outputPath, "_clip") {
		t.Error("expected output path to contain '_clip'")
	}

	// Check arguments contain key flags
	hasInput := false
	hasCopy := false
	hasNoOverwrite := false
	for _, a := range argv {
		if a == "-i" {
			hasInput = true
		}
		if a == "copy" {
			hasCopy = true
		}
		if a == "-n" {
			hasNoOverwrite = true
		}
	}

	if !hasInput {
		t.Error("expected -i flag in arguments")
	}
	if !hasCopy {
		t.Error("expected 'copy' encoder in arguments")
	}
	if !hasNoOverwrite {
		t.Error("expected -n (no overwrite) flag in arguments")
	}
}

func TestBuildClipArgsGPUEncoding(t *testing.T) {
	args := ClipArgs{
		InputFile: "/tmp/test.mp4",
		StartTime: 0,
		EndTime:   60 * time.Second,
		OutputDir: "/tmp/output",
		Encoder:   "h264_amf",
	}

	argv, _ := BuildClipArgs(args)

	hasAMF := false
	hasAudioCopy := false
	for _, a := range argv {
		if a == "h264_amf" {
			hasAMF = true
		}
		if a == "copy" {
			hasAudioCopy = true
		}
	}

	if !hasAMF {
		t.Error("expected h264_amf encoder in arguments")
	}
	if !hasAudioCopy {
		t.Error("expected audio copy flag in GPU arguments")
	}
}

func TestBuildExportArgsWithScale(t *testing.T) {
	args := ExportArgs{
		InputFile: "/tmp/test.mp4",
		OutputDir: "/tmp/output",
		Encoder:   "libx264",
		Width:     1920,
		Height:    1080,
		Bitrate:   "10M",
	}

	argv, _ := BuildExportArgs(args)

	hasVf := false
	hasBitrate := false
	for _, a := range argv {
		if a == "-vf" {
			hasVf = true
		}
		if a == "-b:v" {
			hasBitrate = true
		}
	}

	if !hasVf {
		t.Error("expected -vf flag for scaling")
	}
	if !hasBitrate {
		t.Error("expected -b:v flag for bitrate")
	}
}

func TestBuildExportArgsStreamCopy(t *testing.T) {
	args := ExportArgs{
		InputFile: "/tmp/test.mp4",
		OutputDir: "/tmp/output",
		Encoder:   "copy",
	}

	argv, _ := BuildExportArgs(args)

	hasCopy := false
	for _, a := range argv {
		if a == "copy" {
			hasCopy = true
		}
	}

	if !hasCopy {
		t.Error("expected copy encoder for stream copy")
	}
}

func TestBuildClipArgsTimestampOrder(t *testing.T) {
	args := ClipArgs{
		InputFile: "/tmp/test.mp4",
		StartTime: 5 * time.Second,
		EndTime:   10 * time.Second,
		OutputDir: "/tmp/output",
		Encoder:   "copy",
	}

	argv, _ := BuildClipArgs(args)

	// Verify -ss comes before -to
	ssIndex := -1
	toIndex := -1
	for i, a := range argv {
		if a == "-ss" {
			ssIndex = i
		}
		if a == "-to" {
			toIndex = i
		}
	}

	if ssIndex < 0 {
		t.Error("expected -ss flag")
	}
	if toIndex < 0 {
		t.Error("expected -to flag")
	}
	if ssIndex >= toIndex {
		t.Error("expected -ss before -to")
	}
}

func TestValidateTimestamps(t *testing.T) {
	tests := []struct {
		name    string
		start   time.Duration
		end     time.Duration
		dur     time.Duration
		wantErr bool
	}{
		{"valid", 0, 10 * time.Second, 60 * time.Second, false},
		{"negative start", -1, 10 * time.Second, 60 * time.Second, true},
		{"end before start", 10 * time.Second, 5 * time.Second, 60 * time.Second, true},
		{"equal start end", 10 * time.Second, 10 * time.Second, 60 * time.Second, true},
		{"exceeds duration", 0, 70 * time.Second, 60 * time.Second, true},
		{"zero duration ok", 0, 10 * time.Second, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTimestamps(tt.start, tt.end, tt.dur)
			if tt.wantErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
