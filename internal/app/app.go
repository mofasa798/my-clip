package app

import (
	"context"

	"my-clip/internal/domain"
	"my-clip/internal/source/download"
	"my-clip/internal/source/registry"
	"my-clip/internal/source/resolver"
	"my-clip/internal/system"
)

// Options holds the dependencies for the application service.
type Options struct {
	Logger        *system.Logger
	Config        *system.Config
	Deps          *system.DepResult
	Detector      *system.Detector
	SourceRegistry *registry.Registry
}

// App is the main application service exposed to the frontend via Wails bindings.
type App struct {
	logger   *system.Logger
	config   *system.Config
	deps     *system.DepResult
	detector *system.Detector

	sourceResolver *resolver.Resolver
	downloadSvc    *download.Service
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
	return "0.1.0"
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
// url: the video URL
// streamID: the selected stream format ID
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

// onDownloadProgress is called periodically during download.
func (a *App) onDownloadProgress(p domain.DownloadProgress) {
	a.logger.Debug("Download progress: %.1f%%", p.Percentage)
}
