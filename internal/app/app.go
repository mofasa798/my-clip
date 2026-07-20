package app

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"my-clip/internal/domain"
	"my-clip/internal/media/export"
	"my-clip/internal/media/ffmpeg"
	"my-clip/internal/media/ffprobe"
	"my-clip/internal/media/gpu"
	"my-clip/internal/source/download"
	"my-clip/internal/source/registry"
	"my-clip/internal/source/resolver"
	"my-clip/internal/system"
)

// Options holds the dependencies for the application service.
type Options struct {
	Logger         *system.Logger
	Config         *system.Config
	Deps           *system.DepResult
	Detector       *system.Detector
	SourceRegistry *registry.Registry
	ExportService  *export.Service
	GPUDetector    *gpu.Detector
	FFmpegWrapper  *ffmpeg.Wrapper
	HistoryStore   *system.HistoryStore
}

// App is the main application service exposed to the frontend via Wails bindings.
type App struct {
	logger   *system.Logger
	config   *system.Config
	deps     *system.DepResult
	detector *system.Detector

	sourceResolver *resolver.Resolver
	downloadSvc    *download.Service
	exportSvc      *export.Service
	gpuDetector    *gpu.Detector
	ffmpeg         *ffmpeg.Wrapper
	history        *system.HistoryStore
}

// New creates a new App service.
func New(opts Options) *App {
	svc := download.New(opts.SourceRegistry, opts.Config.OutputDir)

	return &App{
		logger:         opts.Logger,
		config:         opts.Config,
		deps:           opts.Deps,
		detector:       opts.Detector,
		sourceResolver: resolver.New(opts.SourceRegistry),
		downloadSvc:    svc,
		exportSvc:      opts.ExportService,
		gpuDetector:    opts.GPUDetector,
		ffmpeg:         opts.FFmpegWrapper,
		history:        opts.HistoryStore,
	}
}

// Startup is called when the application starts.
func (a *App) Startup() error {
	a.logger.Info("App service starting up")
	return nil
}

// Shutdown is called when the application is shutting down.
func (a *App) Shutdown() error {
	a.logger.Info("App service shutting down")
	return nil
}

// GetVersion returns the current application version.
func (a *App) GetVersion() string {
	return "0.2.0"
}

// GetDependencies returns the current dependency detection results.
func (a *App) GetDependencies() *system.DepResult {
	return a.deps
}

// RefreshDependencies re-checks all dependencies and returns updated results.
func (a *App) RefreshDependencies() *system.DepResult {
	a.deps = a.detector.DetectAll()
	a.logger.Info("Dependencies refreshed")
	return a.deps
}

// GetConfig returns the current application configuration.
func (a *App) GetConfig() *system.Config {
	return a.config
}

// SaveConfig saves new configuration values.
func (a *App) SaveConfig(cfg *system.Config) error {
	a.config = cfg
	if err := system.SaveConfig(cfg); err != nil {
		a.logger.Error("Failed to save config: %v", err)
		return err
	}
	a.logger.Info("Configuration saved")
	return nil
}

// --- Source Layer Methods ---

// SupportedSources returns the list of available video sources.
func (a *App) SupportedSources() []string {
	return a.downloadSvc.SupportedSources()
}

// ResolveSource checks if a URL is supported and returns the source name.
func (a *App) ResolveSource(url string) (string, error) {
	return a.downloadSvc.LookupURL(url)
}

// GetMetadata retrieves video metadata for the given URL.
func (a *App) GetMetadata(url string) (*domain.VideoMetadata, error) {
	ctx := context.Background()
	meta, err := a.downloadSvc.GetMetadata(ctx, url)
	if err != nil {
		a.logger.Error("Failed to get metadata: %v", err)
		return nil, err
	}
	a.logger.Info("Metadata retrieved for %s: %s", meta.Source, meta.Title)
	return meta, nil
}

// StartDownload begins downloading a video.
func (a *App) StartDownload(url, streamID string) (*domain.DownloadResult, error) {
	ctx := context.Background()
	a.logger.Info("Starting download: %s", url)

	result, err := a.downloadSvc.Download(ctx, url, streamID, a.onDownloadProgress)
	if err != nil {
		a.logger.Error("Download failed: %v", err)
		return nil, err
	}

	a.logger.Info("Download completed: %s", result.FilePath)
	return result, nil
}

func (a *App) onDownloadProgress(p domain.DownloadProgress) {
	a.logger.Debug("Download progress: %.1f%%", p.Percentage)
}

// --- Media Layer Methods ---

// ProbeFile reads metadata from a local media file using FFprobe.
func (a *App) ProbeFile(filePath string) (*ffprobe.MediaInfo, error) {
	a.logger.Info("Probing file: %s", filePath)
	info, err := ffprobe.Probe(filePath)
	if err != nil {
		a.logger.Error("FFprobe failed: %v", err)
		return nil, err
	}
	return info, nil
}

// CreateClip extracts a clip from a local media file.
func (a *App) CreateClip(inputFile string, startSeconds, endSeconds float64) error {
	a.logger.Info("Creating clip: %s [%f - %f]", inputFile, startSeconds, endSeconds)

	req := domain.ClipRequest{
		InputFile: inputFile,
		StartTime: time.Duration(startSeconds * float64(time.Second)),
		EndTime:   time.Duration(endSeconds * float64(time.Second)),
		OutputDir: a.config.OutputDir,
	}

	return a.exportSvc.CreateClip(req, a.onExportProgress)
}

// ExportFile exports a media file with the given options.
func (a *App) ExportFile(inputFile, encoder, format string) error {
	a.logger.Info("Exporting file: %s (encoder: %s)", inputFile, encoder)

	req := domain.ExportRequest{
		InputFile: inputFile,
		OutputDir: a.config.OutputDir,
		Encoder:   encoder,
		Format:    format,
	}

	return a.exportSvc.Export(req, a.onExportProgress)
}

// GetGPUInfo returns GPU encoding capabilities.
func (a *App) GetGPUInfo() *gpu.Capabilities {
	return a.gpuDetector.Detect()
}

// GetAvailableEncoders returns the list of available encoders.
func (a *App) GetAvailableEncoders() []domain.EncoderOption {
	return a.exportSvc.AvailableEncoders()
}

func (a *App) onExportProgress(p domain.ExportProgress) {
	a.logger.Debug("Export progress: %.1f%% (%.1f fps, %s)", p.Percentage, p.FPS, p.Speed)
}

// --- History Methods ---

// GetHistory returns all download and export history entries.
func (a *App) GetHistory() ([]domain.HistoryEntry, error) {
	entries, err := a.history.All()
	if err != nil {
		a.logger.Error("Failed to load history: %v", err)
		return nil, err
	}
	return entries, nil
}

// ClearHistory removes all history entries.
func (a *App) ClearHistory() error {
	return a.history.Clear()
}

// --- File Management ---

// GetOutputDir returns the current output directory path.
func (a *App) GetOutputDir() string {
	return a.config.OutputDir
}

// OpenFolder opens the given folder in the file explorer.
func (a *App) OpenFolder(path string) error {
	a.logger.Info("Opening folder: %s", path)
	cmd := exec.Command("explorer", path)
	return cmd.Start()
}

// GetFileInfo returns information about a local file.
func (a *App) GetFileInfo(path string) (map[string]interface{}, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("file not found: %s", path)
	}

	return map[string]interface{}{
		"name":  info.Name(),
		"size":  info.Size(),
		"is_dir": info.IsDir(),
	}, nil
}
