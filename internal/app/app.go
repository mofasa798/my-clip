package app

import (
	"my-clip/internal/system"
)

// Options holds the dependencies for the application service.
type Options struct {
	Logger   *system.Logger
	Config   *system.Config
	Deps     *system.DepResult
	Detector *system.Detector
}

// App is the main application service exposed to the frontend via Wails bindings.
type App struct {
	logger   *system.Logger
	config   *system.Config
	deps     *system.DepResult
	detector *system.Detector
}

// New creates a new App service.
func New(opts Options) *App {
	return &App{
		logger:   opts.Logger,
		config:   opts.Config,
		deps:     opts.Deps,
		detector: opts.Detector,
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
