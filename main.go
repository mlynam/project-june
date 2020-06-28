/*
The main package launches the window, initializes the engine and starts the game
in the startup state.
*/
package main

import (
	"fmt"
	"runtime"

	"github.com/mlynam/project-june/core"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/shaders"
	"github.com/qmuntal/gltf"
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

	doc, _ := gltf.Open("assets/models/cube.gltf")
	fmt.Print(doc)

	core.New(&init).Run()
}
