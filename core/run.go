package core

import "github.com/go-gl/glfw/v3.2/glfw"

// Run is a function that starts the core engine
func (c *Core) Run() {
	defer glfw.Terminate()

	for !c.window.ShouldClose() {
		// TODO: update stuff

		glfw.PollEvents()
	}
}
