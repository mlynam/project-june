package core

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mlynam/project-june/shader"
	"github.com/mlynam/project-june/shared"
)

// Core contains core game engine functionality
type Core struct {
	state                     State
	name                      string
	width                     int
	height                    int
	window                    *glfw.Window
	glProgram                 uint32
	shaders                   []shader.Shader
	secondsSinceLastFPSUpdate float64

	Scene *Scene
}

// Update the core
func (c *Core) Update(ctx *shared.Context) {
	elapsed := ctx.TimeDelta()
	c.secondsSinceLastFPSUpdate += elapsed

	if c.secondsSinceLastFPSUpdate > 1 {
		c.window.SetTitle(fmt.Sprintf("%v - %d fps", c.name, int(1/elapsed)))
		c.secondsSinceLastFPSUpdate = 0
	}
}

// AspectRatio returns the width / height aspect ratio of the window
func (c *Core) AspectRatio() float32 {
	return float32(c.width) / float32(c.height)
}

// OpenGLProgram returns a handle to the graphics program used to render
func (c *Core) OpenGLProgram() uint32 {
	return c.glProgram
}

// Name returns the name of the core
func (c *Core) Name() string {
	return c.name
}
