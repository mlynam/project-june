package manager

import "github.com/mlynam/project-june/engine"

// Provider creates new managers
type Provider struct{}

// New creates a manager
func (p *Provider) New(s engine.Settings) engine.Manager {
	return &Manager{}
}
