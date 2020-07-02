package engine

import "github.com/mlynam/project-june/engine/provider"

type Engine interface {
	Start(string)
}

type Providers struct {
	Graphics provider.Graphics
	Window   provider.Window
	Asset    provider.Asset
	Settings provider.Settings
}

func New(p *Provider) Engine {
	return nil
}
