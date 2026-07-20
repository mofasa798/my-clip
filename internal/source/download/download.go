// Package download orchestrates the download workflow using generic models.
//
// It uses the Source Resolver to find the correct adapter for a URL,
// then delegates the actual download to that adapter. The rest of the
// application never needs to know which source was used.
package download

import (
	"context"
	"fmt"
	"os"

	"my-clip/internal/domain"
	"my-clip/internal/source/registry"
	"my-clip/internal/source/resolver"
)

// Service orchestrates media downloads.
type Service struct {
	resolver  *resolver.Resolver
	registry  *registry.Registry
	outputDir string
}

// New creates a new download Service.
func New(reg *registry.Registry, outputDir string) *Service {
	return &Service{
		resolver:  resolver.New(reg),
		registry:  reg,
		outputDir: outputDir,
	}
}

// LookupURL identifies which source handles the given URL.
// Returns the source name and nil on success, or an error if the URL is unsupported.
func (s *Service) LookupURL(url string) (string, error) {
	adapter, err := s.resolver.Resolve(url)
	if err != nil {
		return "", err
	}
	return adapter.Name(), nil
}

// GetMetadata retrieves metadata for the given URL.
func (s *Service) GetMetadata(ctx context.Context, url string) (*domain.VideoMetadata, error) {
	adapter, err := s.resolver.Resolve(url)
	if err != nil {
		return nil, err
	}
	return adapter.Metadata(ctx, url)
}

// Download downloads media from the given URL using the specified stream.
// progress is an optional callback for download progress updates.
func (s *Service) Download(ctx context.Context, url, streamID string, progress func(domain.DownloadProgress)) (*domain.DownloadResult, error) {
	adapter, err := s.resolver.Resolve(url)
	if err != nil {
		return nil, err
	}

	meta, err := adapter.Metadata(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}

	if err := os.MkdirAll(s.outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	filename := sanitizeFilename(meta.Title) + ".mp4"

	req := domain.DownloadRequest{
		URL:       url,
		StreamID:  streamID,
		OutputDir: s.outputDir,
		Filename:  filename,
	}

	result, err := adapter.Download(ctx, req, progress)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// SupportedSources returns the names of all registered source adapters.
func (s *Service) SupportedSources() []string {
	return s.registry.Names()
}

// sanitizeFilename removes characters that are invalid in filenames.
func sanitizeFilename(name string) string {
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := name
	for _, ch := range invalid {
		result = stringsReplaceAll(result, ch, "_")
	}
	return result
}

func stringsReplaceAll(s, old, new string) string {
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); {
		if i+len(old) <= len(s) && s[i:i+len(old)] == old {
			result = append(result, []byte(new)...)
			i += len(old)
		} else {
			result = append(result, s[i])
			i++
		}
	}
	return string(result)
}

// StorageDir returns the current storage output directory.
func (s *Service) StorageDir() string {
	return s.outputDir
}
