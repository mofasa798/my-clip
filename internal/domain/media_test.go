package domain

import (
	"testing"
	"time"
)

func TestVideoMetadataDefaults(t *testing.T) {
	m := &VideoMetadata{
		Source:   "YouTube",
		Title:    "Test Video",
		Author:   "Test Author",
		Duration: time.Minute * 5,
		URL:      "https://youtube.com/watch?v=test",
	}

	if m.Source != "YouTube" {
		t.Errorf("expected Source YouTube, got %s", m.Source)
	}
	if m.Title != "Test Video" {
		t.Errorf("expected Title 'Test Video', got %s", m.Title)
	}
	if m.Duration != 5*time.Minute {
		t.Errorf("expected Duration 5m, got %v", m.Duration)
	}
	if m.Streams != nil {
		t.Errorf("expected Streams nil, got %v", m.Streams)
	}
}

func TestStreamInfoValidation(t *testing.T) {
	s := StreamInfo{
		ID:        "137",
		Quality:   "1080p",
		Format:    "mp4",
		HasVideo:  true,
		HasAudio:  true,
		Bitrate:   5000,
		Codec:     "h264",
	}

	if !s.HasVideo {
		t.Error("expected HasVideo to be true")
	}
	if !s.HasAudio {
		t.Error("expected HasAudio to be true")
	}
	if s.ID != "137" {
		t.Errorf("expected ID 137, got %s", s.ID)
	}
}

func TestDownloadRequest(t *testing.T) {
	req := DownloadRequest{
		URL:       "https://youtube.com/watch?v=test",
		StreamID:  "best",
		OutputDir: "/tmp/output",
		Filename:  "video.mp4",
	}

	if req.URL == "" {
		t.Error("URL should not be empty")
	}
	if req.OutputDir == "" {
		t.Error("OutputDir should not be empty")
	}
}

func TestClipRequestValidation(t *testing.T) {
	tests := []struct {
		name      string
		start     time.Duration
		end       time.Duration
		expectErr bool
	}{
		{"valid range", 10 * time.Second, 30 * time.Second, false},
		{"zero start", 0, 10 * time.Second, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := ClipRequest{
				InputFile: "/tmp/test.mp4",
				StartTime: tt.start,
				EndTime:   tt.end,
			}
			if req.EndTime <= req.StartTime && !tt.expectErr {
				t.Error("expected valid clip range")
			}
		})
	}
}

func TestExportRequestDefaults(t *testing.T) {
	req := ExportRequest{
		InputFile: "/tmp/test.mp4",
		OutputDir: "/tmp/output",
		Encoder:   "copy",
		Format:    "mp4",
	}

	if req.Encoder != "copy" {
		t.Errorf("expected encoder 'copy', got %s", req.Encoder)
	}
	if req.Format != "mp4" {
		t.Errorf("expected format 'mp4', got %s", req.Format)
	}
}

func TestEncoderOption(t *testing.T) {
	opt := EncoderOption{
		Name:      "Stream Copy",
		Value:     "copy",
		Available: true,
	}

	if !opt.Available {
		t.Error("expected Available to be true")
	}
	if opt.Name != "Stream Copy" {
		t.Errorf("expected name 'Stream Copy', got %s", opt.Name)
	}
}

func TestDownloadResult(t *testing.T) {
	r := &DownloadResult{
		FilePath: "/tmp/output/video.mp4",
		Size:     1048576,
		Duration: 120 * time.Second,
		Format:   "mp4",
	}

	if r.Size != 1048576 {
		t.Errorf("expected size 1048576, got %d", r.Size)
	}
	if r.Duration != 120*time.Second {
		t.Errorf("expected duration 2m, got %v", r.Duration)
	}
}

func TestDownloadProgress(t *testing.T) {
	p := DownloadProgress{
		Percentage:  45.5,
		Speed:       "3.5MiB/s",
		ETA:         "00:17",
		BytesLoaded: 50000000,
		TotalBytes:  110000000,
	}

	if p.Percentage != 45.5 {
		t.Errorf("expected 45.5%%, got %f", p.Percentage)
	}
	if p.TotalBytes <= p.BytesLoaded {
		t.Error("TotalBytes should be greater than BytesLoaded")
	}
}

func TestExportProgress(t *testing.T) {
	p := ExportProgress{
		Percentage: 75.0,
		FPS:        120.5,
		Speed:      "2.5x",
		ETA:        "00:05",
	}

	if p.FPS != 120.5 {
		t.Errorf("expected FPS 120.5, got %f", p.FPS)
	}
}

func TestDefaultPresets(t *testing.T) {
	presets := DefaultPresets()
	if len(presets) != 3 {
		t.Errorf("expected 3 presets, got %d", len(presets))
	}

	expected := map[string]string{
		"Fast (Stream Copy)":  "copy",
		"Balanced (GPU)":      "auto",
		"Maximum Quality":     "libx264",
	}

	for _, p := range presets {
		enc, ok := expected[p.Name]
		if !ok {
			t.Errorf("unexpected preset: %s", p.Name)
			continue
		}
		if p.Encoder != enc {
			t.Errorf("preset %s: expected encoder %s, got %s", p.Name, enc, p.Encoder)
		}
	}
}
