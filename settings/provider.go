package settings

import "github.com/mlynam/project-june/engine"

// Provider for settings
type Provider struct{}

// New settings generated by the provider
func (p *Provider) New() engine.Settings {
	return &Settings{
		name:       "Project June",
		resolution: [2]uint16{600, 600},
	}
}
