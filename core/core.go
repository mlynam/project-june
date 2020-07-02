package core

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// OnCoreAction signature used to perform core actions
type OnCoreAction func(Engine)

// OnLoop signature used when the core is running
type OnLoop func(*Context)

// Core contains core game engine functionality
type core struct {
	state      State
	name       string
	width      int
	height     int
	window     *glfw.Window
	onStartup  OnCoreAction
	onShutdown OnCoreAction
	onLoop     OnLoop
}

// Run the core
func (c *core) Run() {
	if c.onLoop == nil {
		panic("No loop callback defined. Use `core.Builder.UseStartup` to configure the core")
	}

	if c.onStartup != nil {
		c.onStartup(c)
	}

	// Cleanup actions
	defer func() {
		c.state = Shutdown

		if c.onShutdown != nil {
			c.onShutdown(c)
		}

		glfw.Terminate()
	}()

	// Start running
	c.state = Running
	lastFrameTime := glfw.GetTime()

	for !c.window.ShouldClose() {
		// Timing logic
		frameTime := glfw.GetTime()
		elapsed := frameTime - lastFrameTime
		lastFrameTime = frameTime
		context := NewContext(elapsed, c)

		c.onLoop(context)

		// Show frame and listen to platform events
		c.window.SwapBuffers()
		glfw.PollEvents()
	}
}

// AspectRatio returns the width / height aspect ratio of the window
func (c *core) AspectRatio() float32 {
	return float32(c.width) / float32(c.height)
}

// Name returns the name of the core
func (c *core) Name() string {
	return c.name
}
