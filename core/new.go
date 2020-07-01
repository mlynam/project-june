package core

import (
	"github.com/mlynam/project-june/shader"
)

// New is a function that initializes the core
func New(init *Init) *Core {
	c := Core{}

	c.state = Startup
	c.name = init.Name
	c.width = init.Width
	c.height = init.Height
	c.shaders = make([]shader.Shader, 0)
	c.secondsSinceLastFPSUpdate = 1

	return c.initWindow().initGraphics(&init.Graphics)
}
