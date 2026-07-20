// Package registry maintains a registry of available source adapters.
//
// New sources are registered here and the resolver uses this registry
// to find the correct adapter for a given URL.
package registry

import (
	"fmt"
	"sync"

	"my-clip/internal/source"
)

// Registry holds all registered source adapters.
type Registry struct {
	mu      sync.RWMutex
	adapters map[string]source.SourceAdapter
}

// New creates a new empty Registry.
func New() *Registry {
	return &Registry{
		adapters: make(map[string]source.SourceAdapter),
	}
}

// Register adds a source adapter to the registry.
func (r *Registry) Register(adapter source.SourceAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.adapters[adapter.Name()] = adapter
}

// Get returns the adapter with the given name.
func (r *Registry) Get(name string) (source.SourceAdapter, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	adapter, ok := r.adapters[name]
	if !ok {
		return nil, fmt.Errorf("source adapter not found: %s", name)
	}
	return adapter, nil
}

// All returns all registered adapters.
func (r *Registry) All() []source.SourceAdapter {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]source.SourceAdapter, 0, len(r.adapters))
	for _, a := range r.adapters {
		result = append(result, a)
	}
	return result
}

// Names returns the names of all registered adapters.
func (r *Registry) Names() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]string, 0, len(r.adapters))
	for name := range r.adapters {
		result = append(result, name)
	}
	return result
}
