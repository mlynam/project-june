package core

import "github.com/go-gl/glfw/v3.2/glfw"

func (c *Core) initWindow() *Core {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(c.width, c.height, c.name, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	c.window = window

	return c
}
