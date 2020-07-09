package platform

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mlynam/project-june/engine"
)

// NewWindow creates a new window
func (p *Provider) NewWindow(s engine.Settings) engine.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	width := s.Resolution()[0]
	height := s.Resolution()[1]

	window, err := glfw.CreateWindow(int(width), int(height), s.Name(), nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	return window
}
