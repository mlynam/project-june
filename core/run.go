package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mlynam/project-june/shared"
)

// Run is a function that starts the core engine
func (c *Core) Run() {
	defer glfw.Terminate()

	gl.UseProgram(c.glProgram)
	c.Scene.Camera.SetTransformations(c.glProgram)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(100/256.0, 149/256.0, 237/256.0, 1.0)

	lastFrameTime := glfw.GetTime()

	for !c.window.ShouldClose() {
		// Frame setup
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Timing logic
		frameTime := glfw.GetTime()
		elapsed := frameTime - lastFrameTime
		lastFrameTime = frameTime
		context := shared.NewContext(elapsed, c)

		// Update logic
		for _, updatable := range c.Scene.Updatables {
			updatable.Update(context)
		}

		// Render logic
		for _, renderable := range c.Scene.Renderables {
			renderable.Render(c.glProgram)
		}

		// Show frame and listen to platform events
		c.window.SwapBuffers()
		glfw.PollEvents()
	}
}
