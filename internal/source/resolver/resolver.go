// Package resolver matches a video URL to the correct source adapter.
//
// The resolver iterates through all registered adapters and returns
// the first one that matches the URL.
package resolver

import (
	"fmt"

	"my-clip/internal/source"
	"my-clip/internal/source/registry"
)

// Resolver finds the correct source adapter for a given URL.
type Resolver struct {
	registry *registry.Registry
}

// New creates a new Resolver that uses the given registry.
func New(reg *registry.Registry) *Resolver {
	return &Resolver{
		registry: reg,
	}
}

// Resolve finds the adapter that matches the given URL.
// Returns an error if no adapter matches.
func (r *Resolver) Resolve(url string) (source.SourceAdapter, error) {
	for _, adapter := range r.registry.All() {
		if adapter.Match(url) {
			return adapter, nil
		}
	}
	return nil, fmt.Errorf("unsupported video URL: %s", url)
}
