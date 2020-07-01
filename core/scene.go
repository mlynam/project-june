package core

import (
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/shared"
)

// Scene contains all scene data
type Scene struct {
	Renderables []graphics.Renderable
	Updatables  []shared.Updatable
	Camera      *graphics.Camera
}
