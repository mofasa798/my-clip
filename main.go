package main

import (
	"embed"
	"log"

	"my-clip/internal/app"
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

	// Initialize application
	appInstance := app.New(app.Options{
		Logger:     logger,
		Config:     cfg,
		Deps:       deps,
		Detector:   detector,
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
