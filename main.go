/*
The main package launches the window, initializes the engine and starts the game
in the startup state.
*/
package main

func main() {
	e := engine.New()
	e.Start("assets/scene/BoxTaurus.gltf")
}
