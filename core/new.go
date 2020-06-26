package core

import (
	"github.com/mlynam/project-june/shaders"
)

// New is a function that initializes the core
func New(init *Init) *Core {
	c := Core{}

	c.state = Startup
	c.name = init.Name
	c.width = init.Width
	c.height = init.Height
	c.shaders = make([]shaders.Shader, 0)

	return c.initWindow().initGraphics(&init.Graphics)
}
