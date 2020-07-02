package provider

import "github.com/mlynam/project-june/engine"

// Window provider creates a window for the engine to use
type Window interface {
	Create() engine.Window
}
