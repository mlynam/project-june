package settings

import "github.com/mlynam/project-june/engine"

type Provider struct{}

func (p *Provider) New() engine.Settings {
	return &Settings{
		name:       "Project June",
		resolution: [2]uint16{1920, 1080},
	}
}
