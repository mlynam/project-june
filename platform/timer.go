package platform

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mlynam/project-june/engine"
)

// Timer keeps time
type Timer struct{}

// NewTimer returns a Timer
func (p *Provider) NewTimer(s engine.Settings) engine.Timer {
	return &Timer{}
}

// GetTime returns the time in seconds
func (t *Timer) GetTime() float64 {
	return glfw.GetTime()
}
