package main

import (
	"embed"
	"log"

	"my-clip/internal/app"
	"my-clip/internal/media/export"
	"my-clip/internal/media/ffmpeg"
	"my-clip/internal/media/gpu"
	"my-clip/internal/source/kick"
	"my-clip/internal/source/registry"
	"my-clip/internal/source/ytdlp"
	"my-clip/internal/source/youtube"
	"my-clip/internal/system"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Initialize logger
	logger := system.NewLogger()
	logger.Info("Starting My Clip application")

	// Load configuration
	cfg, err := system.LoadConfig()
	if err != nil {
		logger.Warn("Could not load config, using defaults: %v", err)
		cfg = system.DefaultConfig()
	}

	// Detect dependencies
	detector := system.NewDetector()
	deps := detector.DetectAll()
	logger.Info("Dependency detection complete")

	// Initialize yt-dlp wrapper
	yt, err := ytdlp.New()
	if err != nil {
		logger.Warn("yt-dlp not available: %v", err)
	}

	// Initialize source registry and register adapters
	sourceRegistry := registry.New()

	if yt != nil {
		ytAdapter, ytErr := youtube.New(yt)
		if ytErr == nil {
			sourceRegistry.Register(ytAdapter)
			logger.Info("YouTube adapter registered")
		}

		kickAdapter, kickErr := kick.New(yt)
		if kickErr == nil {
			sourceRegistry.Register(kickAdapter)
			logger.Info("Kick adapter registered")
		}
	}

	// Initialize media layer
	gpuDetector := gpu.New()
	gpuCaps := gpuDetector.Detect()
	logger.Info("GPU detection complete: available=%v, preferred=%s", gpuCaps.GPUAvailable, gpuCaps.Preferred)

	ffmpegWrapper, fwErr := ffmpeg.New()
	if fwErr != nil {
		logger.Warn("ffmpeg not available: %v", fwErr)
	}

	var exportSvc *export.Service
	if ffmpegWrapper != nil {
		exportSvc = export.New(ffmpegWrapper, gpuDetector, cfg.OutputDir)
		logger.Info("Export service initialized")
	}

	// Initialize history store
	historyStore := system.NewHistoryStore()
	logger.Info("History store initialized")

	// Initialize preset store
	presetStore := system.NewPresetStore()
	logger.Info("Preset store initialized")

	// Initialize application
	appInstance := app.New(app.Options{
		Logger:         logger,
		Config:         cfg,
		Deps:           deps,
		Detector:       detector,
		SourceRegistry: sourceRegistry,
		ExportService:  exportSvc,
		GPUDetector:    gpuDetector,
		FFmpegWrapper:  ffmpegWrapper,
		HistoryStore:   historyStore,
		PresetStore:    presetStore,
	})

	// Create Wails application
	wailsApp := application.New(application.Options{
		Name:        "My Clip",
		Description: "Multi-Platform Video Clipper",
		Services: []application.Service{
			application.NewService(appInstance),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create main window
	wailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "My Clip",
		Width:  1000,
		Height: 680,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(15, 15, 20),
		URL:              "/",
	})

	logger.Info("Application started successfully")

	// Run application
	err = wailsApp.Run()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Application shutdown complete")
}
