package provider

import "github.com/mlynam/project-june/engine"

// Settings provider gets the engine settings
type Settings interface {
	Get() *engine.Settings
}
