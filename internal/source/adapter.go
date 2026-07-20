// Package source defines the SourceAdapter interface and related types.
//
// All video source implementations (YouTube, Kick, etc.) must implement
// this interface. The rest of the application works only with generic
// domain models and never depends on a specific source.
package source

import (
	"context"

	"my-clip/internal/domain"
)

// SourceAdapter is the interface that every video source must implement.
type SourceAdapter interface {
	// Name returns the display name of the source (e.g. "YouTube").
	Name() string

	// Match returns true if the given URL belongs to this source.
	Match(url string) bool

	// Metadata retrieves video metadata for the given URL.
	Metadata(ctx context.Context, url string) (*domain.VideoMetadata, error)

	// Download downloads media for the given request and returns the result.
	Download(ctx context.Context, req domain.DownloadRequest, progress func(domain.DownloadProgress)) (*domain.DownloadResult, error)
}
