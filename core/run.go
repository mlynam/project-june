package core

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Run is a function that starts the core engine
func (c *Core) Run() {
	defer glfw.Terminate()

	for !c.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(c.glProgram)

		for _, r := range c.Renderables {
			(*r).Render()
		}

		glfw.PollEvents()
		c.window.SwapBuffers()
	}
}
