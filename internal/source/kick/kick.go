// Package kick implements the SourceAdapter interface for Kick.
package kick

import (
	"context"
	"strings"

	"my-clip/internal/domain"
	"my-clip/internal/source"
	"my-clip/internal/source/ytdlp"
)

const sourceName = "Kick"

// Adapter implements source.SourceAdapter for Kick.
type Adapter struct {
	ytdlp *ytdlp.Wrapper
}

// New creates a new Kick adapter.
func New(yt *ytdlp.Wrapper) (*Adapter, error) {
	return &Adapter{ytdlp: yt}, nil
}

// Name returns the display name of this source.
func (a *Adapter) Name() string {
	return sourceName
}

// Match returns true if the URL belongs to Kick.
func (a *Adapter) Match(url string) bool {
	lower := strings.ToLower(url)
	return strings.Contains(lower, "kick.com") ||
		strings.Contains(lower, "kick.com/video")
}

// Metadata retrieves video metadata from Kick.
func (a *Adapter) Metadata(ctx context.Context, url string) (*domain.VideoMetadata, error) {
	return a.ytdlp.ExtractMetadata(ctx, url)
}

// Download downloads a Kick video.
func (a *Adapter) Download(ctx context.Context, req domain.DownloadRequest, progress func(domain.DownloadProgress)) (*domain.DownloadResult, error) {
	return a.ytdlp.Download(ctx, req, progress)
}

// Ensure Adapter implements SourceAdapter.
var _ source.SourceAdapter = (*Adapter)(nil)
