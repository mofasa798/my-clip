// Package export orchestrates clip extraction and media export using FFmpeg.
//
// It provides a high-level API for clipping and exporting while handling
// encoder selection, validation, and progress reporting automatically.
package export

import (
	"fmt"
	"os"

	"my-clip/internal/domain"
	"my-clip/internal/media/ffmpeg"
	"my-clip/internal/media/gpu"
)

// Service provides clip extraction and media export operations.
type Service struct {
	ffmpeg   *ffmpeg.Wrapper
	gpu      *gpu.Detector
	outputDir string
}

// New creates a new export Service.
func New(fw *ffmpeg.Wrapper, gpuDetector *gpu.Detector, outputDir string) *Service {
	return &Service{
		ffmpeg:   fw,
		gpu:      gpuDetector,
		outputDir: outputDir,
	}
}

// CreateClip extracts a clip from a local media file.
// The encoder is selected automatically based on availability:
// stream copy → GPU → CPU.
func (s *Service) CreateClip(req domain.ClipRequest, progress func(domain.ExportProgress)) error {
	// Validate input file
	if _, err := os.Stat(req.InputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file not found: %s", req.InputFile)
	}

	// Validate timestamps
	// Duration is optional for validation; pass 0 to skip duration check
	if err := ffmpeg.ValidateTimestamps(req.StartTime, req.EndTime, 0); err != nil {
		return fmt.Errorf("invalid clip timestamps: %w", err)
	}

	// Select encoder
	encoder := s.selectEncoder(false)

	clipArgs := ffmpeg.ClipArgs{
		InputFile: req.InputFile,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		OutputDir: req.OutputDir,
		Encoder:   encoder,
	}

	return s.ffmpeg.RunClip(clipArgs, progress)
}

// Export processes a media file with the given export options.
func (s *Service) Export(req domain.ExportRequest, progress func(domain.ExportProgress)) error {
	if _, err := os.Stat(req.InputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file not found: %s", req.InputFile)
	}

	encoder := req.Encoder
	if encoder == "" || encoder == "auto" {
		encoder = s.selectEncoder(false)
	}

	exportArgs := ffmpeg.ExportArgs{
		InputFile: req.InputFile,
		OutputDir: req.OutputDir,
		Encoder:   encoder,
		Format:    req.Format,
	}

	return s.ffmpeg.RunExport(exportArgs, 0, progress)
}

// selectEncoder chooses the best available encoder based on strategy.
// If reEncode is false, stream copy is preferred.
func (s *Service) selectEncoder(reEncode bool) string {
	caps := s.gpu.Detect()

	if !reEncode && caps.StreamCopy {
		return "copy"
	}

	if caps.GPUAvailable && caps.Preferred != "" {
		return caps.Preferred
	}

	return "libx264"
}

// GetCapabilities returns the current GPU encoding capabilities.
func (s *Service) GetCapabilities() *gpu.Capabilities {
	return s.gpu.Detect()
}

// GetOutputDir returns the current output directory.
func (s *Service) GetOutputDir() string {
	return s.outputDir
}

// AvailableEncoders returns the list of encoders that can be used.
func (s *Service) AvailableEncoders() []domain.EncoderOption {
	caps := s.gpu.Detect()

	options := []domain.EncoderOption{
		{Name: "Stream Copy", Value: "copy", Available: true},
	}

	for _, enc := range caps.Encoders {
		if enc.Available {
			label := enc.Name
			switch enc.Name {
			case "h264_amf":
				label = "AMD AMF (H.264)"
			case "hevc_amf":
				label = "AMD AMF (HEVC)"
			case "h264_nvenc":
				label = "NVIDIA NVENC (H.264)"
			case "hevc_nvenc":
				label = "NVIDIA NVENC (HEVC)"
			case "h264_qsv":
				label = "Intel QSV (H.264)"
			}
			options = append(options, domain.EncoderOption{
				Name:      label,
				Value:     enc.Name,
				Available: true,
			})
		}
	}

	options = append(options, domain.EncoderOption{
		Name:      "CPU (libx264)",
		Value:     "libx264",
		Available: true,
	})

	return options
}

// StorageInfo holds information about storage directories.
type StorageInfo struct {
	DownloadDir string `json:"download_dir"`
	ExportDir   string `json:"export_dir"`
	FreeBytes   int64  `json:"free_bytes"`
}

// GetStorageInfo returns information about storage directories.
func GetStorageInfo(downloadDir, exportDir string) (*StorageInfo, error) {
	info := &StorageInfo{
		DownloadDir: downloadDir,
		ExportDir:   exportDir,
	}

	// Check download directory
	if err := os.MkdirAll(downloadDir, 0755); err == nil {
		info.FreeBytes = getFreeSpace(downloadDir)
	}

	return info, nil
}

// getFreeSpace returns available disk space in bytes.
func getFreeSpace(path string) int64 {
	// On Windows, we could use syscall.GetDiskFreeSpaceEx
	// For now, return 0 to indicate unknown
	return 0
}
