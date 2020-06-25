/*
The main package launches the window, initializes the engine and starts the game
in the startup state.
*/
package main

import (
	"fmt"

	"github.com/mlynam/project-june/core"
)

func main() {
	core.Init()
	fmt.Println("Hello, world!")
}
