package platform

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mlynam/project-june/engine"
)

// Platform access
type Platform struct{}

// NewPlatform returns a new platform instance
func (p *Provider) NewPlatform(engine.Settings) engine.Platform {
	return &Platform{}
}

// PollEvents waits for events from the platform
func (p *Platform) PollEvents() {
	glfw.PollEvents()
}

// Terminate the platform
func (p *Platform) Terminate() {
	glfw.Terminate()
}
