package provider

import "github.com/mlynam/project-june/engine"

// Graphics provider creates a graphics engine
type Graphics interface {
	Create() engine.Graphics
}
