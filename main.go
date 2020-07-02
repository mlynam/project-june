/*
The main package launches the window, initializes the engine and starts the game
in the startup state.
*/
package main

import (
	"github.com/mlynam/project-june/engine"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/manager"
	"github.com/mlynam/project-june/platform"
	"github.com/mlynam/project-june/settings"
)

func main() {
	e := engine.NewHost(&engine.HostProviders{
		Platform: &platform.Provider{},
		Graphics: &graphics.Provider{},
		Settings: &settings.Provider{},
		Manager:  &manager.Provider{},
	})

	e.Run("BoxTaurus")
}
