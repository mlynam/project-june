/*
The main package launches the window, initializes the engine and starts the game
in the startup state.
*/
package main

import (
	"runtime"

	"github.com/mlynam/project-june/core"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/shaders"
)

func main() {
	runtime.LockOSThread()

	init := core.Init{
		Name:   "Project June",
		Width:  1920,
		Height: 1080,
		Graphics: graphics.Init{
			Shaders: map[shaders.ShaderType]string{
				shaders.Vertex:   "assets/shaders/basic.vert",
				shaders.Fragment: "assets/shaders/basic.frag",
			},
		},
	}

	core.New(&init).Run()
}
