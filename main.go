/*
The main package launches the window, initializes the engine and starts the game
in the startup state.
*/
package main

import (
	"runtime"

	"github.com/mlynam/project-june/core"
	"github.com/mlynam/project-june/game"
	"github.com/mlynam/project-june/graphics/shader"
	"github.com/mlynam/project-june/shared"
)

var (
	world   *game.World
	program *shader.Program
)

func main() {
	runtime.LockOSThread()

	builder := core.NewBuilder()

	builder.UseLoop(func(c *shared.Context) {
		world.Update(c)
		world.Render(program)
	})

	builder.Build().Run()
}
