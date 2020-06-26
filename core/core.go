package core

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mlynam/project-june/shaders"
)

// Core contains core game engine functionality
type Core struct {
	state     State
	name      string
	width     int
	height    int
	window    *glfw.Window
	glProgram uint32
	shaders   []shaders.Shader
}
