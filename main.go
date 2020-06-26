/*
The main package launches the window, initializes the engine and starts the game
in the startup state.
*/
package main

import (
	"runtime"

	"github.com/mlynam/project-june/core"
)

func main() {
	runtime.LockOSThread()

	init := core.Init{
		Name:   "Project June",
		Width:  1920,
		Height: 1080,
	}

	core.New(&init).Run()
}
